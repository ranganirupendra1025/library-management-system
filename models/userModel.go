package models

import (
	"time"
)

type User struct {
	user_id               int
	username              string
	password              string
	age                   int
	is_approved           bool
	email_address         string
	is_admin              bool
	subscription_id       int
	subscription_end_date time.Time
}
