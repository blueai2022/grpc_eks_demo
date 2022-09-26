package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/blueai2022/appsubmission/api"
	db "github.com/blueai2022/appsubmission/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secretpwd@localhost:5432/app_submission?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
	fmt.Println("LifeAI app submission api has started")
}
