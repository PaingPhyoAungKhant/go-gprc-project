package main

import (
	"database/sql"
	api "go-gprc-project/api"
	db "go-gprc-project/db/sqlc"
	"go-gprc-project/util"
	"log"

	_ "github.com/lib/pq"
)

//
// const (
// 	dbSource   = "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable"
// 	driverName = "postgres"
// 	address    = "0.0.0.0:8081"
// )

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot load db", err)
	}

	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot start server: %w", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal(err)
	}
}
