package controller

import (
	"database/sql"
	"encoding/json"
	"lms/models"
	"lms/service"
	"net/http"
)

//IssueBook issues a book to user
func IssueBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction models.Transaction
		err := json.NewDecoder(r.Body).Decode(&transaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//Call the service  method to issue the book
		err = service.IssueBook(transaction.UserID, transaction.BookID, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//Respond with success
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Book Issued Successfully"})
	}
}

//ReturnBook returns a book issued by user
func ReturnBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction models.Transaction
		err := json.NewDecoder(r.Body).Decode(&transaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//Call the service  method to return the book
		err = service.IssueBook(transaction.UserID, transaction.BookID, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//Respond with success
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Book Returned Successfully"})

	}
}
