package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/yuhengfdada/go-bank/db"
	"github.com/yuhengfdada/go-bank/token"
	"github.com/yuhengfdada/go-bank/util"
)

type Server struct {
	store          db.Store
	router         *gin.Engine
	tokenGenerator token.TokenGenerator
}

func NewServer(store db.Store) *Server {
	router := gin.Default()
	tokenGen, err := token.NewPasetoTokenGenerator(util.GetRandomString(32))
	if err != nil {
		log.Fatal(err)
	}
	server := Server{
		store:          store,
		router:         router,
		tokenGenerator: tokenGen,
	}
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("validCurrency", validCurrency)
	}
	server.router.POST("/accounts", server.createAccount) // POST
	server.router.GET("/accounts/:id", server.getAccount) // GET with uri params
	server.router.GET("/accounts", server.listAccount)    // GET with query params (?key=value)
	server.router.POST("/transfer", server.createTransfer)
	server.router.POST("/user", server.createUser)
	server.router.POST("/login", server.loginUser)

	return &server
}

func (server *Server) Start() {
	server.router.Run()
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
