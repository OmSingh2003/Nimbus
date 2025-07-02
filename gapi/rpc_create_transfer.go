package gapi

import (
	"context"

	db "github.com/OmSingh2003/vaultguard-api/db/sqlc"
	"github.com/OmSingh2003/vaultguard-api/pb"
	"github.com/OmSingh2003/vaultguard-api/val"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateTransfer(ctx context.Context, req *pb.CreateTransferRequest) (*pb.CreateTransferResponse, error) {
	authPayload, err := server.getAuthPayload(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", err)
	}

	err = validateCreateTransferRequest(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err)
	}

	// Verify that the from_account belongs to the authenticated user
	fromAccount, err := server.store.GetAccount(ctx, req.GetFromAccountId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get from account: %s", err)
	}

	if fromAccount.Owner != authPayload.Username {
		return nil, status.Errorf(codes.PermissionDenied, "from account doesn't belong to the authenticated user")
	}

	// Verify that the to_account exists
	toAccount, err := server.store.GetAccount(ctx, req.GetToAccountId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get to account: %s", err)
	}

	// Check if currency matches
	if fromAccount.Currency != req.GetCurrency() {
		return nil, status.Errorf(codes.InvalidArgument, "from account currency mismatch: %s vs %s", fromAccount.Currency, req.GetCurrency())
	}

	if toAccount.Currency != req.GetCurrency() {
		return nil, status.Errorf(codes.InvalidArgument, "to account currency mismatch: %s vs %s", toAccount.Currency, req.GetCurrency())
	}

	arg := db.TransferTxParams{
		FromAccountID: req.GetFromAccountId(),
		ToAccountID:   req.GetToAccountId(),
		Amount:        req.GetAmount(),
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create transfer: %s", err)
	}

	rsp := &pb.CreateTransferResponse{
		Transfer: convertTransfer(result.Transfer),
	}

	return rsp, nil
}

func validateCreateTransferRequest(req *pb.CreateTransferRequest) error {
	if req.GetFromAccountId() <= 0 {
		return status.Errorf(codes.InvalidArgument, "from_account_id must be greater than 0")
	}

	if req.GetToAccountId() <= 0 {
		return status.Errorf(codes.InvalidArgument, "to_account_id must be greater than 0")
	}

	if req.GetAmount() <= 0 {
		return status.Errorf(codes.InvalidArgument, "amount must be greater than 0")
	}

	if err := val.ValidateCurrency(req.GetCurrency()); err != nil {
		return status.Errorf(codes.InvalidArgument, "invalid currency: %s", err.Error())
	}

	return nil
}

func convertTransfer(transfer db.Transfer) *pb.Transfer {
	return &pb.Transfer{
		Id:            transfer.ID,
		FromAccountId: transfer.FromAccountID,
		ToAccountId:   transfer.ToAccountID,
		Amount:        transfer.Amount,
		Currency:      "", // Note: Transfer table doesn't store currency directly
		CreatedAt:     transfer.CreatedAt.String(),
	}
}
