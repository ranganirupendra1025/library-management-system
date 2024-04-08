package controller

import (
	"encoding/json"
	"lms/models"
	"lms/service"
	"net/http"
)

//BookController handles HTTP requests related to books
type BookController struct {
	Service service.BookService
}

//IssueBook issues a book to user
func (c *BookController) IssueBook(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Call the service  method to issue the book
	err = c.Service.IssueBook(transaction.UserID, transaction.BookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Respond with success
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Book Issued Successfully"})
}

//ReturnBook returns a book issued by user
func (c *BookController) ReturnBook(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Call the service  method to return the book
	err = c.Service.IssueBook(transaction.UserID, transaction.BookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Respond with success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Book Returned Successfully"})

}