package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"lms/models"
	"lms/service"
	"net/http"
	"strconv"
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
func GetSingleUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//id, err := strconv.Atoi(r.URL.Query().Get("id"))
		idstr := r.URL.Path[len("/getuser/"):]
		id, err := strconv.Atoi(idstr)

		//eventId:=ctx.Param("id")
		if err != nil {
			http.Error(w, "Invalid User id", http.StatusBadRequest)
			return

		}
      fmt.Println(id)
		user, err := service.GetUser(db, id)
		if err != nil {
			http.Error(w, "No user found", http.StatusBadRequest)
			return

		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
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

		err = service.RegisterUser(db, newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "User Added Successfully")

	}
}

func DeletingUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idstr := r.URL.Path[len("/deleteuser/"):]
		id, err := strconv.Atoi(idstr)

		if err != nil {
			http.Error(w, "No id specified", http.StatusBadRequest)
			return

		}
		user, err := service.GetUser(db, id)
		if err != nil {
			http.Error(w, "Invalid user id", http.StatusInternalServerError)
			return
		}
		var userAuth models.UserAuth
		err = json.NewDecoder(r.Body).Decode(&userAuth)
		if err != nil {
			http.Error(w, "Authentication failed/Give Admin id", http.StatusBadRequest)
			return
		}
		users, err := service.Authenticate(userAuth.Adminid, db)
		if err != nil {
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return

		}
		if !users.Isadmin {
			http.Error(w, "Only Admin have an Access", http.StatusBadRequest)
			return
		}

		err = service.DeleteUser(db, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User Deleted Successfully")

	}
}

func LoginUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.Login
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		err = service.Login(db, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User login Successfully")

	}
}
