package handler

import (
	"net/http"
	"samproj/controllers"
)

func defineSubscriptionHandlers() {
	http.HandleFunc("/users/v1/signup", controllers.SignUp)
}
