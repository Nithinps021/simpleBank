package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/nithinps021/simplebank/db/sqlc"
	"github.com/nithinps021/simplebank/util"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6" `
	FullName string `json:"fullname" binding:"required"`
	EmialID  string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	Username            string    `json:"username"`
	FullName            string    `json:"full_name"`
	EmialID             string    `json:"emial_id"`
	PasswordLastChanged time.Time `json:"password_last_changed"`
	CreatedAt           time.Time `json:"created_at"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorReponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	arg := db.CreateUserParams{
		Username:     req.Username,
		FullName:     req.FullName,
		HashPassword: hashedPassword,
		EmialID:      req.EmialID,
	}
	account, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqError, ok := err.(*pq.Error); ok {
			switch pqError.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, pqError)
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorReponse(err))
		return
	}
	response:=createUserResponse{
		Username: account.Username,
		FullName: account.FullName,
		EmialID: account.EmialID,
		PasswordLastChanged: account.PasswordLastChanged,
		CreatedAt: account.CreatedAt,
	}
	ctx.JSON(http.StatusOK, response)
}
