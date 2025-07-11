package main

import (
	"database/sql"
	"log"

	db "github.com/DhawalYerekar/Users-App/DB/sqlc"
	"github.com/DhawalYerekar/Users-App/api"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://myfoodsapp:mysecretpassword@localhost:5432/foodsapp?sslmode=disable"
)

func main() {
	// connect to db

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Failed to connect Database")
	}

	//initialise the store
	store := db.NewStore(conn)

	//initialize server
	server, err := api.NewServer(*store)
	if err != nil {
		log.Fatal("Cannot create server")
	}

	//run server on port
	err = server.Start(":8080")
	if err != nil {
		log.Fatal("Failed to start server", err)
	}

}
