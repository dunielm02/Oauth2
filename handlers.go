package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	userdata := &User{}

	err := json.NewDecoder(r.Body).Decode(userdata)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userdata.Password_hash), bcrypt.DefaultCost)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = DB.QueryRow(
		"INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id",
		userdata.Username,
		hashedPass,
	).Scan(&userdata.ID)

	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	log.Println("Successful: User: ", userdata.Username)
	w.WriteHeader(200)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var usr User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	var hashed_password string

	err = DB.QueryRow("SELECT id, password_hash FROM users WHERE users.username = $1", usr.Username).Scan(&usr.ID, &hashed_password)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(usr.Password_hash))
	if err != nil {
		ErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	token := rand.Text()

	hashedToken := sha256.New()
	hashedToken.Write([]byte(token))

	_, err = DB.Exec(
		"INSERT INTO sessions (username_id, token) VALUES ($1, $2) RETURNING id",
		usr.ID,
		hashedToken.Sum(nil),
	)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	jsonToken, err := json.Marshal(token)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err)
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(jsonToken)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("alive"))
}
