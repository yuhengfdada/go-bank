package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/yuhengfdada/go-bank/db"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	router := gin.Default()
	server := Server{
		store:  store,
		router: router,
	}
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("validCurrency", validCurrency)
	}
	server.router.POST("/accounts", server.createAccount) // POST
	server.router.GET("/accounts/:id", server.getAccount) // GET with uri params
	server.router.GET("/accounts", server.listAccount)    // GET with query params (?key=value)
	server.router.POST("/transfer", server.createTransfer)

	return &server
}

func (server *Server) Start() {
	server.router.Run()
}
