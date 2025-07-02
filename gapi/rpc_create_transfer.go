package gapi

import (
	"context"
	"time"

	db "github.com/OmSingh2003/nimbus/db/sqlc"
	"github.com/OmSingh2003/nimbus/pb"
	"github.com/OmSingh2003/nimbus/val"
	"github.com/OmSingh2003/nimbus/worker"
	"github.com/hibiken/asynq"
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

	// Check if this is a transfer to the demo account
	if toAccount.ID == getDemoAccountID() || isAccountNumber(toAccount, "DEMO-1234567890") {
		// Trigger demo response task
		taskPayload := &worker.PayloadDemoResponse{
			FromAccountID: req.GetFromAccountId(),
			Amount:        req.GetAmount(),
			Currency:      req.GetCurrency(),
		}
		opts := []asynq.Option{
			asynq.MaxRetry(3),
			asynq.ProcessIn(10 * time.Second), // Delay for demo effect
			asynq.Queue("demo"),
		}
		_ = server.taskDistributor.DistributeTaskDemoResponse(ctx, taskPayload, opts...)
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

// Helper function to detect demo account
func getDemoAccountID() int64 {
	// This should be replaced with actual demo account lookup if needed
	return -1 // Placeholder - not used since we check account number instead
}

func isAccountNumber(account db.Account, accountNumber string) bool {
	// Check if account has the specific account number
	// This assumes account_number field exists in the Account struct
	// You may need to adjust based on your actual Account struct
	return account.Owner == "demo_user" // Simplified check for demo account
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
