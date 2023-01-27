package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/nithinps021/simplebank/db/sqlc"
)

type transactionRequest struct {
	FromAccountId int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountId   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,min=1"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransaction(ctx *gin.Context) {
	var req transactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorReponse(err))
		return
	}
	if !server.validateAccount(ctx,req.FromAccountId,req.Currency){
		return
	} 
	if !server.validateAccount(ctx,req.ToAccountId,req.Currency){
		return
	} 

	arg := db.AddTransferParams{
		FromAccountID: req.FromAccountId,
		ToAccountID:   req.ToAccountId,
		Amount:        req.Amount,
	}
	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorReponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validateAccount(ctx *gin.Context, accountId int64, currency string) bool {
	account, err := server.store.GetAccount(ctx, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorReponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorReponse(err))
		return false
	}
	if account.Currency!=currency{
		err:=fmt.Errorf("account [%d] has different currency : [%s] vs [%s]",accountId,account.Currency,currency)
		ctx.JSON(http.StatusBadRequest,errorReponse(err))
		return false
	}
	return true
}
