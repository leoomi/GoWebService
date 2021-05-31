package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Pool *sql.DB

func Init() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	Pool, err = sql.Open("postgres", connStr)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name", err)
	}

	Pool.SetConnMaxLifetime(0)
	Pool.SetMaxIdleConns(3)
	Pool.SetMaxOpenConns(3)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = Pool.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func Close() {
	Pool.Close()
}
