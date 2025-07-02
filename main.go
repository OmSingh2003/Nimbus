package main

import (
	"context"
	"database/sql"
	"net"
	"net/http"
	"os"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"

	"github.com/OmSingh2003/vaultguard-api/api"
	db "github.com/OmSingh2003/vaultguard-api/db/sqlc"
	_ "github.com/OmSingh2003/vaultguard-api/doc/statik"
	"github.com/OmSingh2003/vaultguard-api/gapi"
	"github.com/OmSingh2003/vaultguard-api/mail"
	"github.com/OmSingh2003/vaultguard-api/pb"
	"github.com/OmSingh2003/vaultguard-api/util"
	"github.com/OmSingh2003/vaultguard-api/worker"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot load configurations")
	}

	// Use PORT environment variable for Render deployment
	if port := os.Getenv("PORT"); port != "" {
		config.HTTPServerAddress = "0.0.0.0:" + port
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to the database")
	}

	store := db.NewStore(conn)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

	go runTaskProcessor(&config, redisOpt, store)
	// Run both HTTP Gateway and gRPC servers concurrently
	go runGatewayServer(config, store, taskDistributor)
	runGrpcServer(config, store, taskDistributor)
}

func runGatewayServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create gRPC server")
	}

	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterVaultguardAPIHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register handle server")
	}
	mux := http.NewServeMux()
	// Add CORS middleware
	corsHandler := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
			if r.Method == "OPTIONS" {
				return
			}
			h.ServeHTTP(w, r)
		})
	}
	mux.Handle("/", corsHandler(grpcMux))

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create statik fs")
	}
	swagger_handler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swagger_handler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create gateway listener")
	}

	log.Info().Msgf("Start HTTP gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot start HTTP gateway server")
	}
}

func runGrpcServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create gRPC server")
	}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			gapi.GrpcLogger,
			server.AuthorizationInterceptor,
		),
	)
	pb.RegisterVaultguardAPIServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create gRPC listener")
	}

	log.Info().Msgf("start gRPC server at [%s]", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot start gRPC server")
	}
}

func runGinServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := api.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create HTTP server")
	}

	log.Printf("Start HTTP server at %s", config.HTTPServerAddress)
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot start HTTP server")
	}
}

func runTaskProcessor(config *util.Config, redisOpt asynq.RedisClientOpt, store db.Store) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer, config)
	log.Info().Msg("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}
}
