package api

import (
	"fmt"

	"github.com/blueai2022/appsubmission/config"
	db "github.com/blueai2022/appsubmission/db/sqlc"
	"github.com/blueai2022/appsubmission/token"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     *config.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config *config.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	//one or more routes here
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	router.GET("/users/:username", server.getUser)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
