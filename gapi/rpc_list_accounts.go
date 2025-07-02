package gapi

import (
	"context"
	"errors"

	db "github.com/OmSingh2003/vaultguard-api/db/sqlc"
	"github.com/OmSingh2003/vaultguard-api/pb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListAccounts(ctx context.Context, req *pb.ListAccountsRequest) (*pb.ListAccountsResponse, error) {
	authPayload, err := server.getAuthPayload(ctx)
	if err != nil {
		return nil, err
	}

	violations := validateListAccountsRequest(req)
	if violations != nil {
		return nil, InvalidArgumentError(violations)
	}

	arg := db.ListAccountsParams{
		Owner:  authPayload.Username,
		Limit:  req.GetPageSize(),
		Offset: (req.GetPageId() - 1) * req.GetPageSize(),
	}

	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list accounts: %s", err)
	}

	pbAccounts := make([]*pb.Account, len(accounts))
	for i, account := range accounts {
		pbAccounts[i] = convertAccount(account)
	}

	rsp := &pb.ListAccountsResponse{
		Accounts: pbAccounts,
	}
	return rsp, nil
}

func validateListAccountsRequest(req *pb.ListAccountsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.GetPageId() <= 0 {
		violations = append(violations, fieldViolation("page_id", errors.New("must be greater than 0")))
	}

	if req.GetPageSize() <= 0 {
		violations = append(violations, fieldViolation("page_size", errors.New("must be greater than 0")))
	} else if req.GetPageSize() > 100 {
		violations = append(violations, fieldViolation("page_size", errors.New("must not be greater than 100")))
	}

	return violations
}
