package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	//"fmt"
	"io"
	"io/ioutil"
	"lms/models"
	"lms/service"
	"net/http"
	"strconv"
)

func Addingbooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//ADMIN ACCESS
		body,err:=io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
         var userAuth models.UserAuth
		// err = json.NewDecoder(r.Body).Decode(&userAuth)
		err=json.Unmarshal(body,&userAuth)
		if err != nil {
			http.Error(w, "Authentication failed/give username", http.StatusInternalServerError)
			return
		}
	
		
		
		user,err:=service.Authenticate(userAuth,db)
		if err != nil {
			http.Error(w, "No user with that name", http.StatusInternalServerError)
			return
		}
      if(!user.Isadmin){
		http.Error(w, "Only Admin user can access", http.StatusBadRequest)
			return

	  }
	 // r.Body.Seek(0,0)
	 r.Body=ioutil.NopCloser(bytes.NewBuffer(body))
		var book models.Book
		err= json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = service.Addbooks(db, book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		//fmt.Fprintf(w, "Book was added")*/
		json.NewEncoder(w).Encode("Book added Successsfully")
	}
}
func GetAllBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := service.GettingBooks(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(books)
	}

}
func GetSingleBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//id, err := strconv.Atoi(r.URL.Query().Get("id"))
		idstr:=r.URL.Path[len("/getbook/"):]
		id,err:=strconv.Atoi(idstr)

		//eventId:=ctx.Param("id")
		if err != nil {
			http.Error(w, "Invalid Book id", http.StatusBadRequest)
			return

		}
		
		books, err := service.GetBook(db,id)
		if err != nil {
			http.Error(w, "No Books found", http.StatusBadRequest)
			return

		}
       w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(books)
	}

}


func UpdateStockBooks(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		//id,err:=strconv.ParseInt(r.URL.Query().Get("id"),10,64)
		idstr:=r.URL.Path[len("/update/"):]
		id,err:=strconv.Atoi(idstr)

		if err!=nil{
			http.Error(w, "No id specified", http.StatusBadRequest)
			return

		}
		_,err=service.GetBook(db,id)
		if err!=nil{
			http.Error(w, "No book associated with this ID", http.StatusBadRequest)
			return

		}
		body,err:=io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error in reading request body", http.StatusInternalServerError)
			return
		}
		var userAuth models.UserAuth
		// err = json.NewDecoder(r.Body).Decode(&userAuth)
		err=json.Unmarshal(body,&userAuth)
		if err != nil {
			http.Error(w, "Authentication failed/Give username", http.StatusUnauthorized)
			return
		}
		user,err:=service.Authenticate(userAuth,db)
		if err != nil {
			http.Error(w, "No user with this name", http.StatusUnauthorized)
			return
		}
      if(!user.Isadmin){
		http.Error(w, "Only Admin have an access", http.StatusUnauthorized)
			return

	  }

	 // r.Body.Seek(0,0)
	 r.Body=ioutil.NopCloser(bytes.NewBuffer(body))
		
		var updatedBooks models.Book
		err = json.NewDecoder(r.Body).Decode(&updatedBooks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//updatedBooks.Bookid=id
		err=service.UpdateStock(db,updatedBooks,id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		//fmt.Fprint(w,"Data updated successfully",id)
		json.NewEncoder(w).Encode("Data updated successfully")
	}
	
		
		
}

func Deletingbooks(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		idstr:=r.URL.Path[len("/delete/"):]
		id,err:=strconv.Atoi(idstr)

		if err!=nil{
			http.Error(w, "No id specified", http.StatusBadRequest)
			return

		}
		
		book,err:=service.GetBook(db,id)
		if err!=nil{
			http.Error(w, "No book associated with this ID", http.StatusBadRequest)
			return

		}
		var userAuth models.UserAuth
		err=json.NewDecoder(r.Body).Decode(&userAuth)
		if err!=nil{
			http.Error(w,"Authentication failed/Give username", http.StatusBadRequest)
			return

		}
		user,err:=service.Authenticate(userAuth,db)
		if err != nil {
			http.Error(w, "No user with this name", http.StatusUnauthorized)
			return
		}
      if(!user.Isadmin){
		http.Error(w, "Only Admin have an access", http.StatusUnauthorized)
			return

	  }
		err=service.Delete(book,db)
		if err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Book Deleted Successfully")

	}
}
	

