package controller

import (
	"database/sql"
	"encoding/json"
	"lms/service"
	"net/http"
)

func GetAllSubscription(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := service.GetAllSubscription(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
		}
		json.NewEncoder(w).Encode(users)
	}
}
