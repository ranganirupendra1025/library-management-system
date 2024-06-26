package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"lms/models"
	"lms/service"
	"lms/utils"
	"math"
	"net/http"
	"strconv"
	"time"
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
		user, err := service.GetUser(db, transaction.UserId)
		if err != nil {
			http.Error(w, "No user available with this id", http.StatusInternalServerError)
			return
		}
		//Call the service  method to issue the book
		err = service.IssueBook(transaction.UserId, transaction.BookId, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		msg := ""
		returnTime := time.Now().AddDate(0, 0, utils.BookBorrowPeriodInDays)
		if user.Subdate.Before(time.Now()) {
			msg = fmt.Sprintf(" No active subscription for the user. Please collect Rs.%d", utils.BookBorrowPeriodInDays*utils.BookCostPerDay)
		} else if returnTime.Before(user.Subdate) {
			msg = " Active subscription available. User need not to pay any amount. "
		} else {
			msg = fmt.Sprintf(" Active subscription ends by %s. So please collect Rs.%f",
				user.Subdate, math.Ceil(returnTime.Sub(user.Subdate).Hours()/24)*utils.BookCostPerDay)
		}

		//Respond with success
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Book Issued Successfully." + msg)
	}
}

func GetUserBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//id, err := strconv.Atoi(r.URL.Query().Get("userId"))
		idstr := r.URL.Path[len("/userbooks/"):]
		id, err := strconv.Atoi(idstr)
		// give userid in request url not user_book_transaction table id

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
		//id, err := strconv.Atoi(r.URL.Query().Get("userId"))
		idstr := r.URL.Path[len("/userduebooks/"):]
		id, err := strconv.Atoi(idstr)

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
		//id, err := strconv.Atoi(r.URL.Query().Get("userId"))
		idstr := r.URL.Path[len("/useroverduebooks/"):]
		id, err := strconv.Atoi(idstr)

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

		msg := ""

		userBook, err := service.GetUserBook(transaction.UserId, transaction.BookId, db)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if userBook.ReturnDate.Before(time.Now()) {
			msg = fmt.Sprintf(" Book should have returned on or before %s. Please collect the fine amount of Rs. %f as book returned late", userBook.ReturnDate, math.Ceil(time.Since(userBook.ReturnDate).Hours()/24)*utils.FineAmountPerDay)
		}

		//Call the service  method to return the book
		err = service.ReturnBook(transaction.UserId, transaction.BookId, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//Respond with success
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Book Returned Successfully." + msg)

	}
}
