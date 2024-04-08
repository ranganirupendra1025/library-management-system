package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	//"io"
	"lms/models"
	"lms/service"
	"net/http"
	"time"
	"strconv"
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

func SubscribeUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idstr := r.URL.Path[len("/lms/subscribe/"):]
		id, err := strconv.Atoi(idstr)

		if err != nil {
			http.Error(w, "No id specified", http.StatusBadRequest)
			return

		}
		/*body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}*///validationSuccess := service.ValidateAdmin(r, w, db, subDto)
		// !validationSuccess {
			//return
		//}

		var subDto models.SubscriptionDto
		err=json.NewDecoder(r.Body).Decode(&subDto)
		if err != nil {
			http.Error(w, "Authentication failed/Give Admin Id", http.StatusBadRequest)
			return
		}

		validationSuccess := service.ValidateAdmin(r, w, db, subDto)
	  if !validationSuccess {
			return
		}
		/*if err != nil {
			http.Error(w, "Authentication failed/Give username", http.StatusBadRequest)
			return

		}
		user, err := service.Authenticate(subDto.UserId, db)
		if err != nil {
			http.Error(w, "No user with this name", http.StatusUnauthorized)
			return
		}
		if !user.Isadmin {
			http.Error(w, "Only Admin have an access", http.StatusUnauthorized)
			return

		}*/

		subscription, err := service.GetSubscriptionById(db, subDto.SubscriptionId)
		if err != nil {
			http.Error(w, "No subscription available with this id", http.StatusInternalServerError)
			return
		}
		var userSub models.User
		userSub.Id = id
		userSub.Subid = subDto.SubscriptionId
		userSub.Subdate = <-time.After(time.Duration(subscription.Duration*24*60*60*10 ^ 9))
		//fmt.Println(userSub)
		err = service.UpdateSubscription(db, userSub)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		//fmt.Fprintf(w, "Book was added")*/
		//json.NewEncoder(w).Encode("User Subscrobed")

		json.NewEncoder(w).Encode(fmt.Sprintf("User Subscribed Successsfully. Subscription end date %v", userSub.Subdate))

	}
}
