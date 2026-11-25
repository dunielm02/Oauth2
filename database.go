package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func NewDb() *sql.DB {
	_ = godotenv.Load()

	dsn := os.Getenv("POSTGRES_URL")

	fmt.Println(dsn)

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = conn.Ping()
	if err != nil {
		panic("Not connected to database")
	}

	return conn
}
