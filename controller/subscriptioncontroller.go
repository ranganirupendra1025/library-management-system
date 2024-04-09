package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"

	//"io"
	"lms/models"
	"lms/service"
	"net/http"
	"strconv"
	"time"
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

		var subDto models.SubscriptionDto
		err = json.NewDecoder(r.Body).Decode(&subDto)
		if err != nil {
			http.Error(w, "Authentication failed/Give Admin Id", http.StatusBadRequest)
			return
		}

		validationSuccess := service.ValidateAdmin(r, w, db, subDto)
		if !validationSuccess {
			return
		}

		subscription, err := service.GetSubscriptionById(db, subDto.SubscriptionId)
		if err != nil {
			http.Error(w, "No subscription available with this id", http.StatusInternalServerError)
			return
		}
		var user models.User
		user.Id = id
		user.Subid = subDto.SubscriptionId
		time.Now()
		user.Subdate = <-time.After(time.Duration(subscription.Duration*24*60*60*10 ^ 9))
		//fmt.Println(userSub)
		err = service.UpdateSubscription(db, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		//fmt.Fprintf(w, "Book was added")*/
		//json.NewEncoder(w).Encode("User Subscrobed")

		json.NewEncoder(w).Encode(fmt.Sprintf("User Subscribed Successsfully. Subscription end date %v", user.Subdate))

	}
}
