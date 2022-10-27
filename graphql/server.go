package graphql

import (
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/blueai2022/appsubmission/config"
	db "github.com/blueai2022/appsubmission/db/sqlc"
)

type Server struct {
	config *config.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(config *config.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}
	server.setupRouter()

	return server, nil

}

// Defining the Graphql handler
func graphqlHandler(config *config.Config, store db.Store) gin.HandlerFunc {
	cfg := Config{
		Resolvers: &Resolver{
			Config: config,
			Store:  store,
		},
	}

	h := handler.NewDefaultServer(NewExecutableSchema(cfg))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() error {
	// Setting up Gin
	router := gin.Default()
	router.POST("/query", graphqlHandler(server.config, server.store))
	router.GET("/", playgroundHandler())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8000"},
		AllowMethods:     []string{"POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	server.router = router

	return nil
}
