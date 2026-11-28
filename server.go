package main

import (
	"log"
	"net/http"
)

func NewServer() *http.Server {
	mux := http.NewServeMux()

	var s *http.Server = &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.HandleFunc("POST /register", registerHandler)
	mux.HandleFunc("POST /login", loginHandler)
	mux.HandleFunc("GET /health", healthCheckHandler)

	return s
}

func ErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	log.Println("StatusCode: ", statusCode, " Error: ", err)
}
