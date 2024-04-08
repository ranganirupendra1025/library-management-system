package models

import "time"

type Transaction struct {
	UserID       int       `json:"user_id"`
	BookID       int       `json:"book_id"`
	Timestamp    time.Time `json:"timestamp"`
	ReturnStatus bool      `json:"return_status"`
}
