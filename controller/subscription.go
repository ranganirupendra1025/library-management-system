package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lms/models"
	"lms/service"
	"net/http"
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
		validationSuccess, body := service.ValidateAdmin(r, w, db)
		if !validationSuccess {
			return
		}

		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		var userSub models.User
		err := json.NewDecoder(r.Body).Decode(&userSub)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		subscription, err := service.GetSubscriptionById(db, userSub.Subid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userSub.Subdate = <-time.After(time.Duration(subscription.Duration*24*60*60*10 ^ 9))
		err = service.UpdateSubscription(db, userSub)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		//fmt.Fprintf(w, "Book was added")*/

		json.NewEncoder(w).Encode(fmt.Sprintf("User Subscribed Successsfully. Subscription end date %v", userSub.Subdate))

	}
}
