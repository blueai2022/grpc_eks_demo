package grpcapi

import (
	"fmt"

	"github.com/blueai2022/appsubmission/config"
	db "github.com/blueai2022/appsubmission/db/sqlc"
	"github.com/blueai2022/appsubmission/pb"
	"github.com/blueai2022/appsubmission/token"
)

type Server struct {
	pb.UnimplementedLifeAIServer
	config     *config.Config
	store      db.Store
	tokenMaker token.Maker
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

	return server, nil
}
