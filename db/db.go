package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	//replace with username password and database
	//connStr := "postgres://username:password@localhost/database?sslmode=verify-full"
	connStr := "postgres://postgres:root@localhost/postgres?sslmode=disable"
	DB, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	err = DB.Ping()
	if err != nil {

		return nil, err
	}

	fmt.Println("Successfully connected to database")

	return DB, nil
}
