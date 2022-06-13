package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/yuhengfdada/go-bank/db/code"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	router := gin.Default()
	server := Server{
		store:  store,
		router: router,
	}
	server.router.POST("/accounts", server.createAccount) // POST
	server.router.GET("/accounts/:id", server.getAccount) // GET with uri params
	server.router.GET("/accounts", server.listAccount)    // GET with query params (?key=value)
	return &server
}

func (server *Server) Start() {
	server.router.Run()
}