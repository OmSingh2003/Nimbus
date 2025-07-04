package api 

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/OmSingh2003/nimbus/db/sqlc"
	"github.com/OmSingh2003/nimbus/token"
)


type transferRequest struct {
	FromAccountID    int64 `json:"from_account_id" binding:"required,min=1"` 
	ToAccountID      int64 `json:"to_account_id" binding:"required,min=1"`
	Amount           int64 `json:"amount" binding:"required,gt=0"` 
	Currency         string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}
  fromAccount,valid := server.validAccount(ctx,req.FromAccountID,req.Currency)
 if !valid {
		return 
	}

  authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if fromAccount.Owner != authPayload.Username {
		err := errors.New("from account does not belong to authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return 
	}
	_,valid = server.validAccount(ctx,req.ToAccountID,req.Currency)
	if !valid{
		return
	}
	arg := db.TransferTxParams{
		FromAccountID :req.FromAccountID,
		ToAccountID: req.ToAccountID,
		Amount: req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}

	ctx.JSON(http.StatusOK, result)
}

type listTransfersRequest struct {
	AccountID int64 `form:"account_id" binding:"required,min=1"`
	PageID    int32 `form:"page_id" binding:"required,min=1"`
	PageSize  int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listTransfers(ctx *gin.Context) {
	var req listTransfersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Verify account ownership
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	account, err := server.store.GetAccount(ctx, req.AccountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if account.Owner != authPayload.Username {
		err := errors.New("account does not belong to authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.ListTransfersParams{
		FromAccountID: req.AccountID,
		ToAccountID:   req.AccountID,
		Limit:        req.PageSize,
		Offset:       (req.PageID - 1) * req.PageSize,
	}

	transfers, err := server.store.ListTransfers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transfers)
}

func (server *Server) validAccount (ctx *gin.Context,accountID int64 ,  currency string) (db.Account ,bool) {
	account, err := server.store.GetAccount(ctx,accountID)
	if err != nil {	
		  if err == sql.ErrNoRows {
			 ctx.JSON(http.StatusNotFound, errorResponse(err))
			return account ,false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account ,false
	}
		if account.Currency != currency {
			err := fmt.Errorf("account[%d] currency mismatch : %s vs %s", accountID,account.Currency,currency)
			ctx.JSON(http.StatusBadRequest,errorResponse(err))
			return account ,false 
}
	return account , true 
}
