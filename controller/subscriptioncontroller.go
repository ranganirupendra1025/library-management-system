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
			http.Error(w, "User detail is not specified", http.StatusBadRequest)
			return
		}

		var subDto models.SubscriptionDto
		err = json.NewDecoder(r.Body).Decode(&subDto)
		if err != nil {
			http.Error(w, "Please give Admin Id", http.StatusBadRequest)
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

		user, err := service.GetUser(db, id)
		if err != nil {
			http.Error(w, "No user available with this id", http.StatusInternalServerError)
			return
		}
		user.Subid = subDto.SubscriptionId
		if user.Subdate.IsZero() || user.Subdate.Before(time.Now()) {
			user.Subdate = time.Now().AddDate(0, 0, subscription.Duration)
		} else {
			user.Subdate = user.Subdate.AddDate(0, 0, subscription.Duration)
		}
		//fmt.Println(userSub)
		err = service.UpdateSubscription(db, *user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		//fmt.Fprintf(w, "Book was added")*/
		//json.NewEncoder(w).Encode("User Subscrobed")

		json.NewEncoder(w).Encode(fmt.Sprintf("User Subscribed Successfully. Subscription end date %v.\nPlease collect Rs.%d", user.Subdate, subscription.Cost))

	}
}
