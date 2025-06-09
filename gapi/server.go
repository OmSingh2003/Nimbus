package gapi

import (
	db "github.com/OmSingh2003/vaultguard-api/db/sqlc"
	"github.com/OmSingh2003/vaultguard-api/pb"
	"github.com/OmSingh2003/vaultguard-api/token"
	"github.com/OmSingh2003/vaultguard-api/util"
)

// Server serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedVaultguardAPIServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
