package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/OmSingh2003/vaultguard-api/api"
	db "github.com/OmSingh2003/vaultguard-api/db/sqlc"
	"github.com/OmSingh2003/vaultguard-api/gapi"
	"github.com/OmSingh2003/vaultguard-api/pb"
	"github.com/OmSingh2003/vaultguard-api/util"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	_ "github.com/OmSingh2003/vaultguard-api/doc/statik"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configurations:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}

	store := db.NewStore(conn)

	// Run both HTTP Gateway and gRPC servers concurrently
	go runGatewayServer(config, store)
	runGrpcServer(config, store)
}

func runGatewayServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create gRPC server:", err)
	}

	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterVaultguardAPIHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register handle server")
	}
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal("cannot create statik fs ")
	}
	swagger_handler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swagger_handler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Cannot create gateway listener:", err)
	}

	log.Printf("Start HTTP gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("Cannot start HTTP gateway server:", err)
	}
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create gRPC server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterVaultguardAPIServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Cannot create gRPC listener:", err)
	}

	log.Printf("Start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start gRPC server:", err)
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create HTTP server:", err)
	}

	log.Printf("Start HTTP server at %s", config.HTTPServerAddress)
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Cannot start HTTP server:", err)
	}
}
