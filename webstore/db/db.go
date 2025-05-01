package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	// connection := "user dbname password host sslmode"
	//TODO: This has to move to an environment.
	connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5051, "postgres", "qwerty", "store")

	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		panic(err.Error())
	}

	return db
}
