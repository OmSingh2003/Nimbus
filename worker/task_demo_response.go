package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	db "github.com/OmSingh2003/nimbus/db/sqlc"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskDemoResponse = "task:demo_response"

type PayloadDemoResponse struct {
	FromAccountID int64 `json:"from_account_id"`
	Amount        int64 `json:"amount"`
	Currency      string `json:"currency"`
}

func (distributor *RedisTaskDistributor) DistributeTaskDemoResponse(
	ctx context.Context,
	payload *PayloadDemoResponse,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskDemoResponse, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskDemoResponse(ctx context.Context, task *asynq.Task) error {
	var payload PayloadDemoResponse
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	// Get the demo account by listing accounts for demo_user
	demoAccounts, err := processor.store.ListAccounts(ctx, db.ListAccountsParams{
		Owner:  "demo_user",
		Limit:  1,
		Offset: 0,
	})
	if err != nil {
		return fmt.Errorf("failed to get demo accounts: %w", err)
	}
	if len(demoAccounts) == 0 {
		return fmt.Errorf("no demo account found")
	}
	demoAccount := demoAccounts[0]

	// Get the original sender account
	fromAccount, err := processor.store.GetAccount(ctx, payload.FromAccountID)
	if err != nil {
		return fmt.Errorf("failed to get from account: %w", err)
	}

	// Create demo response transactions
	// Response 1: Send back 50% of the original amount
	response1Amount := payload.Amount / 2
	if response1Amount > 0 {
		_, err = processor.store.TransferTx(ctx, db.TransferTxParams{
			FromAccountID: demoAccount.ID,
			ToAccountID:   payload.FromAccountID,
			Amount:        response1Amount,
		})
		if err != nil {
			log.Error().Err(err).Msg("failed to create demo response 1")
		} else {
			log.Info().Int64("amount", response1Amount).Str("currency", payload.Currency).
				Msg("demo response 1 sent successfully")
		}
	}

	// Wait 5 seconds before second response
	time.Sleep(5 * time.Second)

	// Response 2: Send back 25% of the original amount
	response2Amount := payload.Amount / 4
	if response2Amount > 0 {
		_, err = processor.store.TransferTx(ctx, db.TransferTxParams{
			FromAccountID: demoAccount.ID,
			ToAccountID:   payload.FromAccountID,
			Amount:        response2Amount,
		})
		if err != nil {
			log.Error().Err(err).Msg("failed to create demo response 2")
		} else {
			log.Info().Int64("amount", response2Amount).Str("currency", payload.Currency).
				Msg("demo response 2 sent successfully")
		}
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("from_account", fromAccount.Owner).Msg("processed demo response task")
	return nil
}
