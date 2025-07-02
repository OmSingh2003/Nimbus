package gapi

import (
	db "github.com/OmSingh2003/nimbus/db/sqlc"
	"github.com/OmSingh2003/nimbus/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertAccount(account db.Account) *pb.Account {
	return &pb.Account{
		Id:            account.ID,
		Owner:         account.Owner,
		Balance:       account.Balance,
		Currency:      account.Currency,
		CreatedAt:     timestamppb.New(account.CreatedAt),
		AccountNumber: account.AccountNumber.String,
	}
}
