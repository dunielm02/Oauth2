package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func NewServer(db *sql.DB) *http.Server {
	mux := http.NewServeMux()

	var s *http.Server = &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	mux.HandleFunc("POST /register", func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)

		userdata := &User{}

		err := dec.Decode(userdata)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		hashedPass, err := bcrypt.GenerateFromPassword([]byte(userdata.Password_hash), bcrypt.DefaultCost)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return 
		}

		err = db.QueryRowContext(context.Background(),
		"INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id",
		userdata.Username,
		hashedPass,
		).Scan(&userdata.ID)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(200)

	})

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("alive"))
	})

	return s
}