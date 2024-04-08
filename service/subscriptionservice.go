package service

import (
	"database/sql"
	"lms/models"
	"net/http"
)

//I have added refernce code for getting user from databse  u guys can change as per the requirement
func GetAllSubscription(db *sql.DB) ([]models.Subscription, error) {
	var subscriptions []models.Subscription
	//rows, err := db.Query("SELECT id, name, extract(epoch from duration)*10^9, cost from subscription;")
	rows, err := db.Query("SELECT * FROM subscription")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var subscription models.Subscription
		err := rows.Scan(&subscription.Id, &subscription.Name, &subscription.Duration, &subscription.Cost)
		if err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, subscription)

	}
	return subscriptions, nil

}

func GetSubscriptionById(db *sql.DB, id int) (models.Subscription, error) {

	query := "SELECT * FROM subscription where id=$1"
	row := db.QueryRow(query, id)

	var subscription models.Subscription
	err := row.Scan(&subscription.Id, &subscription.Name, &subscription.Duration, &subscription.Cost)
	if err != nil {
		return models.Subscription{}, err
	}

	return subscription, nil

}

func ValidateAdmin(r *http.Request, w http.ResponseWriter, db *sql.DB, subDto models.SubscriptionDto) bool {

	user, err := Authenticate(subDto.AdminId, db)
	if err != nil {
		http.Error(w, "No user with that name", http.StatusInternalServerError)
		return false
	}
	if !user.Isadmin {
		http.Error(w, "Only Admin user can access", http.StatusBadRequest)
		return false

	}
	return true
}
