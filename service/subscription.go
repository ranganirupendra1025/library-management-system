package service

import (
	"database/sql"
	"lms/models"
)

//I have added refernce code for getting user from databse  u guys can change as per the requirement
func GetAllSubscription(db *sql.DB) ([]models.Subscription, error) {
	var subscriptions []models.Subscription
	//rows, err := db.Query("SELECT id, name, extract(epoch from duration)*10^9, cost from subscription;")
	rows,err:=db.Query("SELECT * FROM subscription")
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

