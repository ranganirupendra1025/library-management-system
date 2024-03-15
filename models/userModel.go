package models

import (
	"time"
)

type User struct {
	UserId              int
	Username            string
	Password            string
	Age                 int
	EmailAddress        string
	IsAdmin             bool
	SubscriptionId      int
	SubscriptionEndDate time.Time
}
