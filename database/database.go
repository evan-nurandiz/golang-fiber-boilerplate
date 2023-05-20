package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func SetupDatabaseConnection() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres",
		"postgres", "go_fiber_boiler_plate")

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")

	return db
}
