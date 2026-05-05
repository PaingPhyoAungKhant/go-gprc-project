package api

import (
	"fmt"
	db "go-gprc-project/db/sqlc"
	token "go-gprc-project/token"
	"go-gprc-project/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
	config     util.Config
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	maker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		tokenMaker: maker,
		config:     config,
	}

	server.SetupRoutes()
	return server, nil
}

func (server *Server) SetupRoutes() {
	router := gin.Default()

	router.POST("/users", server.createUser)

	router.POST("/login", server.loginUser)
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)

	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
