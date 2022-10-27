package main

import (
	"context"
	"database/sql"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	// "github.com/99designs/gqlgen/example/federation/reviews/graph"
	// "github.com/99designs/gqlgen/example/starwars/generated"
	// "github.com/99designs/gqlgen/graphql/handler"
	// "github.com/99designs/gqlgen/graphql/playground"

	"github.com/blueai2022/appsubmission/api"
	"github.com/blueai2022/appsubmission/config"
	db "github.com/blueai2022/appsubmission/db/sqlc"
	_ "github.com/blueai2022/appsubmission/doc/statik"
	"github.com/blueai2022/appsubmission/graphql"
	"github.com/blueai2022/appsubmission/grpcapi"
	"github.com/blueai2022/appsubmission/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	devEnvironment = "development"
)

func main() {
	config, err := config.Load(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	//development: pretty print
	if strings.ToLower(config.Environment) == devEnvironment {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	store := db.NewStore(conn)

	go runGatewayServer(&config, store)
	go runGraphQLServer(&config, store)
	runGrpcServer(&config, store)
	// runGinServer(&config, store)
}

func runGraphQLServer(config *config.Config, store db.Store) {
	// server, err :=
	server, err := graphql.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create GraphQL server")
	}

	log.Info().Msgf("starting GraphQL server at %s", config.GraphQLServerAddress)
	err = server.Start(config.GraphQLServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}

func runGrpcServer(config *config.Config, store db.Store) {
	apiServer, err := grpcapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create api server for gRPC")
	}

	grpcLogger := grpc.UnaryInterceptor(grpcapi.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterLifeAIServer(grpcServer, apiServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener for gRPC")
	}

	log.Info().Msgf("starting gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}

}

func runGatewayServer(config *config.Config, store db.Store) {
	apiServer, err := grpcapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create api server for gRPC")
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	grpcMux := runtime.NewServeMux(jsonOption)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterLifeAIHandlerServer(ctx, grpcMux, apiServer)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register handler server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	// fs := http.FileServer(http.Dir("./doc/swagger"))
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create statik fs")
	}

	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener for HTTP gateway")
	}

	log.Info().Msgf("starting HTTP gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start HTTP gateway server")
	}

}

func runGinServer(config *config.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
