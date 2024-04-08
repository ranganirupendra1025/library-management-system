package models

import "time"

type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	Password string
	Isadmin  bool
	Subid    int
	Subdate  time.Time
}
