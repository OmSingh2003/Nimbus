package gapi

import (
	"context"

	db "github.com/OmSingh2003/nimbus/db/sqlc"
	"github.com/OmSingh2003/nimbus/pb"
	"github.com/OmSingh2003/nimbus/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListTransfers(ctx context.Context, req *pb.ListTransfersRequest) (*pb.ListTransfersResponse, error) {
	authPayload, err := server.getAuthPayload(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	violations := validateListTransfersRequest(req)
	if violations != nil {
		return nil, InvalidArgumentError(violations)
	}

	// Get user's accounts to filter transfers
	accounts, err := server.store.ListAccounts(ctx, db.ListAccountsParams{
		Owner:  authPayload.Username,
		Limit:  100, // Get all user accounts
		Offset: 0,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user accounts: %s", err)
	}

	if len(accounts) == 0 {
		return &pb.ListTransfersResponse{
			Transfers: []*pb.Transfer{},
		}, nil
	}

	// Use the first account for the query (the query will find transfers for any account owned by the user)
	accountId := accounts[0].ID

	arg := db.ListTransfersParams{
		FromAccountID: accountId,
		ToAccountID:   accountId,
		Limit:         req.GetPageSize(),
		Offset:        (req.GetPageId() - 1) * req.GetPageSize(),
	}

	transfers, err := server.store.ListTransfers(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list transfers: %s", err)
	}

	// Filter transfers to only include those involving user's accounts
	var userTransfers []db.Transfer
	accountIds := make(map[int64]bool)
	for _, account := range accounts {
		accountIds[account.ID] = true
	}

	for _, transfer := range transfers {
		if accountIds[transfer.FromAccountID] || accountIds[transfer.ToAccountID] {
			userTransfers = append(userTransfers, transfer)
		}
	}

	// Convert to protobuf
	var pbTransfers []*pb.Transfer
	for _, transfer := range userTransfers {
		pbTransfers = append(pbTransfers, convertTransferWithDetails(transfer, accounts))
	}

	rsp := &pb.ListTransfersResponse{
		Transfers: pbTransfers,
	}

	return rsp, nil
}

func validateListTransfersRequest(req *pb.ListTransfersRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidatePageId(req.GetPageId()); err != nil {
		violations = append(violations, fieldViolation("page_id", err))
	}

	if err := val.ValidatePageSize(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("page_size", err))
	}

	return violations
}

func convertTransferWithDetails(transfer db.Transfer, accounts []db.Account) *pb.Transfer {
	// Find account details for better display
	var fromAccount, toAccount *db.Account
	for i := range accounts {
		if accounts[i].ID == transfer.FromAccountID {
			fromAccount = &accounts[i]
		}
		if accounts[i].ID == transfer.ToAccountID {
			toAccount = &accounts[i]
		}
	}

	// Determine currency (use from_account currency if available)
	currency := "USD" // default
	if fromAccount != nil {
		currency = fromAccount.Currency
	} else if toAccount != nil {
		currency = toAccount.Currency
	}

	return &pb.Transfer{
		Id:            transfer.ID,
		FromAccountId: transfer.FromAccountID,
		ToAccountId:   transfer.ToAccountID,
		Amount:        transfer.Amount,
		Currency:      currency,
		CreatedAt:     transfer.CreatedAt.String(),
	}
}
