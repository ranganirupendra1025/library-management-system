package controller

import (
	"database/sql"
	"encoding/json"
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
