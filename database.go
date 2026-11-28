package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func NewDb() {
	_ = godotenv.Load()

	dsn := os.Getenv("POSTGRES_URL")

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = conn.Ping()
	if err != nil {
		panic("Not connected to database")
	}

	DB = conn
}
