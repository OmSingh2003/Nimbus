package gapi

import (
	db "github.com/OmSingh2003/nimbus/db/sqlc"
	"github.com/OmSingh2003/nimbus/pb"
	"github.com/OmSingh2003/nimbus/token"
	"github.com/OmSingh2003/nimbus/util"
	"github.com/OmSingh2003/nimbus/worker"
)

// Server serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedVaultguardAPIServer
	config         util.Config
	store          db.Store
	tokenMaker     token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
