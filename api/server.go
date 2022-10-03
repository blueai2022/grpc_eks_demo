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

	err = server.setupRouter()
	if err != nil {
		return nil, fmt.Errorf("cannot set up router: %w", err)
	}

	return server, nil
}

func (server *Server) setupRouter() error {
	router := gin.Default()

	//one or more routes here
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker, server.store))

	// g, err := ginproxy.NewGinProxy("http://api.lifeai.us")
	// if err != nil {
	// 	return err
	// }

	authRoutes.POST("/healthai/icd10", setupProxy(server.config.ProxyTargetServer))
	// router.POST("/backend/healthai/icd10", g.Handler)
	// authRoutes.GET("/healthai/icd10", server.recognizeICD)
	authRoutes.GET("/users/:username", server.getUser)

	server.router = router

	return nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
