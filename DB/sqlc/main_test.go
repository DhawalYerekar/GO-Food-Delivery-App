package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://myfoodsapp:mysecretpassword@localhost:5432/foodsapp?sslmode=disable"
)

var testDB *sql.DB
var testQueries *Queries

func TestMain(m *testing.M) {
	var err error

	// 1. Open the DB connection
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("❌ cannot connect to database:", err)
	}

	// // 2. Ping the database to ensure connection is valid
	// err = testDB.Ping()
	// if err != nil {
	// 	log.Fatal("❌ cannot ping database:", err)
	// }

	// 3. Initialize SQLC-generated Queries struct
	testQueries = New(testDB)

	// 4. Run all tests
	os.Exit(m.Run())
}
