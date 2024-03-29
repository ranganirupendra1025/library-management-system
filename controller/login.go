package controller

import (
	"database/sql"
	"encoding/json"
	"lms/models"
	"lms/service"
	"net/http"
)

type UserController struct {
	DB *sql.DB // Database connection
}

func NewUserController(db *sql.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	authenticated, err := service.AuthenticateUser(uc.DB, user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if authenticated {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid username or password"))
	}
}
