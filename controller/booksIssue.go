package controller

import (
	"database/sql"
	"encoding/json"
	"lms/models"
	"lms/service"
	"net/http"
	"strconv"
)

//IssueBook issues a book to user
func IssueBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction models.UserBookTransaction
		err := json.NewDecoder(r.Body).Decode(&transaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//Call the service  method to issue the book
		err = service.IssueBook(transaction.UserId, transaction.BookId, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//Respond with success
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Book Issued Successfully"})
	}
}

func GetUserBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("userId"))
		if err != nil {
			http.Error(w, "Please mention user id", http.StatusBadRequest)
			return
		}
		userBooks, err := service.GetUserBooks(id, db)
		if err != nil {
			http.Error(w, "No records found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(userBooks)
	}
}

func GetUserDueBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("userId"))
		if err != nil {
			http.Error(w, "Please mention user id", http.StatusBadRequest)
			return
		}
		userBooks, err := service.GetUserPendingBooks(id, db)
		if err != nil {
			http.Error(w, "No records found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(userBooks)
	}
}

func GetUserOverdueBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("userId"))
		if err != nil {
			http.Error(w, "Please mention user id", http.StatusBadRequest)
			return
		}
		userBooks, err := service.GetUserOverdueBooks(id, db)
		if err != nil {
			http.Error(w, "No records found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(userBooks)
	}
}

//ReturnBook returns a book issued by user
func ReturnBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction models.UserBookTransaction
		err := json.NewDecoder(r.Body).Decode(&transaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//Call the service  method to return the book
		err = service.IssueBook(transaction.UserId, transaction.BookId, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//Respond with success
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Book Returned Successfully"})

	}
}
