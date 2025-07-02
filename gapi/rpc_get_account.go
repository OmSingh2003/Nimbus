package gapi

import (
	"context"
	"database/sql"
	"errors"

	"github.com/OmSingh2003/vaultguard-api/pb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	authPayload, err := server.getAuthPayload(ctx)
	if err != nil {
		return nil, err
	}

	violations := validateGetAccountRequest(req)
	if violations != nil {
		return nil, InvalidArgumentError(violations)
	}

	account, err := server.store.GetAccount(ctx, req.GetId())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "account not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to find account: %s", err)
	}

	if account.Owner != authPayload.Username {
		return nil, status.Errorf(codes.PermissionDenied, "account doesn't belong to the authenticated user")
	}

	rsp := &pb.GetAccountResponse{
		Account: convertAccount(account),
	}
	return rsp, nil
}

func validateGetAccountRequest(req *pb.GetAccountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.GetId() <= 0 {
		violations = append(violations, fieldViolation("id", errors.New("must be greater than 0")))
	}

	return violations
}
