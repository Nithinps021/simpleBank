package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/nithinps021/simplebank/db/sqlc"
)


type createAccountRequest struct{
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR" `
}

func (server *Server) createAccount(ctx *gin.Context){
	 var req createAccountRequest;
	 if err:=ctx.ShouldBindJSON(&req); err!=nil{
		ctx.JSON(http.StatusBadRequest,errorReponse(err))
		return
	 }
	 arg:=db.CreateAccountParams{
		 Owner: req.Owner,
		 Currency: req.Currency,
		 Balance: 0,
	 }
	 account,err:= server.store.CreateAccount(ctx,arg)
	 if(err!=nil){
		ctx.JSON(http.StatusInternalServerError,errorReponse(err))
		return
	 }
	 ctx.JSON(http.StatusOK,account)
}

type getAccountRequest struct{
	Id int64  `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context){
	var params getAccountRequest
	if err:=ctx.ShouldBindUri(&params); err!=nil{
		ctx.JSON(http.StatusBadRequest,errorReponse(err))
		return
	}
	acccount,err:=server.store.GetAccount(ctx,params.Id)
	if err!=nil{
		if err==sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound,errorReponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError,errorReponse(err))
		return
	}
	ctx.JSON(http.StatusOK,acccount)
}

type listAccountsParams struct{
	PageId int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
func (server *Server) listAccounts(ctx *gin.Context){
	var req listAccountsParams
	if err:=ctx.ShouldBindQuery(&req); err!=nil{
		ctx.JSON(http.StatusInternalServerError,errorReponse(err))
		return
	}
	args:=db.ListAccountsParams{
		Limit: req.PageSize,
		Offset: (req.PageId-1)*req.PageSize,
	}
	fmt.Println(req.PageId,"page_id",req.PageSize," pageSize")
	accounts,err:=server.store.ListAccounts(ctx,args)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,errorReponse(err))
		return
	}
	ctx.JSON(http.StatusOK,accounts)
}

type deleteAccountParams struct{
	Id int64 `json:"id" binding:"required"`
}

func (server *Server) deleteAccount(ctx *gin.Context){
	var req deleteAccountParams
	if err:=ctx.ShouldBindJSON(&req);err!=nil{
		ctx.JSON(http.StatusBadRequest,errorReponse(err))
		return
	}
	err:=server.store.DeleteAccount(ctx,req.Id)
	if err!=nil{
		if err==sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound,errorReponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError,errorReponse(err))
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message":"Account deleted"})
}