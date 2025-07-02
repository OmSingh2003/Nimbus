package gapi

import (
	"context"
	"database/sql"

	db "github.com/OmSingh2003/vaultguard-api/db/sqlc"
	"github.com/OmSingh2003/vaultguard-api/pb"
	"github.com/OmSingh2003/vaultguard-api/util"
	"github.com/OmSingh2003/vaultguard-api/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	authPayload, err := server.getAuthPayload(ctx)
	if err != nil {
		return nil, err
	}

	violations := validateCreateAccountRequest(req)
	if violations != nil {
		return nil, InvalidArgumentError(violations)
	}

	arg := db.CreateAccountParams{
		Owner:         authPayload.Username,
		Currency:      req.GetCurrency(),
		Balance:       10000, // $100 free credit for testing
		AccountNumber: sql.NullString{String: util.RandomAccountNumber(), Valid: true},
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create account: %s", err)
	}

	rsp := &pb.CreateAccountResponse{
		Account: convertAccount(account),
	}
	return rsp, nil
}

func validateCreateAccountRequest(req *pb.CreateAccountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateCurrency(req.GetCurrency()); err != nil {
		violations = append(violations, fieldViolation("currency", err))
	}

	return violations
}

