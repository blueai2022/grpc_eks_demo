package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/blueai2022/appsubmission/api"
	"github.com/blueai2022/appsubmission/config"
	db "github.com/blueai2022/appsubmission/db/sqlc"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.Load(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(&config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
	fmt.Println("LifeAI app submission api has started")
}
