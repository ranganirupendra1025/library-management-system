package controller

import (
	"database/sql"
	"encoding/json"
	"lms/models"
	"lms/service"
	"net/http"
)

func GetAllUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := service.GettingUsers(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(users)
	}

}


func RegisterUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser models.User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newUser.Subid = 1
		err = service.RegisterUser(db, newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}
