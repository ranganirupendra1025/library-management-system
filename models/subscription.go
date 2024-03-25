package models

import "time"

type Subscription struct {
	Id       int
	Name     string
	Duration time.Duration
	Cost     string
}
