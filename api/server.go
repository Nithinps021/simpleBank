package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/nithinps021/simplebank/db/sqlc"
)

type Server struct{
	store *db.Store
	router *gin.Engine
}

func NewServer (store *db.Store) *Server{
	server:= &Server{store: store}
	router:=gin.Default()
	
	if v,ok:=binding.Validator.Engine().(*validator.Validate);ok{
		v.RegisterValidation("currency",validateCurrency)
	}
	router.POST("/user",server.createUser)
	router.POST("/accounts",server.createAccount)
	router.GET("/accounts/:id",server.getAccount)
	router.GET("/accounts",server.listAccounts)
	router.DELETE("/accounts",server.deleteAccount)
	router.POST("/transfer",server.createTransaction)

	server.router= router
	return server
}

func (server *Server) Start(address string) error{
	return server.router.Run(address)
}

func errorReponse(err error) gin.H{
	return gin.H{"error":err.Error()}
}