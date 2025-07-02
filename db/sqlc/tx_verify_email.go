package db

import (
	"context"
	"database/sql"
)

type VerifyEmailTxParams struct {
	EmailId    int64
	SecretCode string
}

type VerifyEmailTxResult struct {
	User           User
	VerifyEmail    VerifyEmail
	WelcomeAccount Account
}

func (store *SQLStore) VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error) {
	var result VerifyEmailTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.VerifyEmail, err = q.UpdateVerifyEmail(ctx, UpdateVerifyEmailParams{
			ID:         arg.EmailId,
			SecretCode: arg.SecretCode,
		})
		if err != nil {
			return err
		}

		result.User, err = q.UpdateUser(ctx, UpdateUserParams{
			Username: result.VerifyEmail.Username,
			IsEmailVerified: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
		})
		if err != nil {
			return err
		}

		// Create welcome account with $100 USD
		result.WelcomeAccount, err = q.CreateAccount(ctx, CreateAccountParams{
			Owner:    result.VerifyEmail.Username,
			Balance:  10000, // $100.00 in cents
			Currency: "USD",
		})
		return err
	})

	return result, err
}
