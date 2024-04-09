package main

import (
	//"database/sql"
	"fmt"
	"io/ioutil"
	"lms/controller"
	"lms/db"
	"log"
	"net/http"
)

func main() {
	//connecting to databse
	DB, err := db.InitDB()
	if err != nil {
		log.Fatal("error in connecting")
	}
	defer DB.Close()

	//executing the script file
	script, err := ioutil.ReadFile("db/init.sql")
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
	}
	_, err = DB.Exec(string(script))

	if err != nil {
		log.Fatal(err)
	}

	//user module APIs
	http.HandleFunc("/registerusers", controller.RegisterUser(DB))
	http.HandleFunc("/users", controller.GetAllUsers(DB))
	http.HandleFunc("/getuser/", controller.GetSingleUser(DB))
	http.HandleFunc("/deleteuser/", controller.DeletingUser(DB))
	http.HandleFunc("/loginuser", controller.LoginUser(DB))

	//subscription module APIs
	http.HandleFunc("/lms/subscriptions", controller.GetAllSubscription(DB))
	http.HandleFunc("/lms/subscribe/", controller.SubscribeUser(DB))

	//book module APIs
	http.HandleFunc("/addbooks", controller.Addingbooks(DB))
	http.HandleFunc("/books", controller.GetAllBooks(DB))
	http.HandleFunc("/getbook/", controller.GetSingleBook(DB))
	http.HandleFunc("/update/", controller.UpdateStockBooks(DB))
	http.HandleFunc("/delete/", controller.Deletingbooks(DB))

	//issue and return  module APIs
	http.HandleFunc("/issuebooks", controller.IssueBook(DB))
	http.HandleFunc("/returnbooks", controller.ReturnBook(DB))

	//server setup
	err = http.ListenAndServe(":7111", nil)
	if err != nil {
		fmt.Println(err)
	}

}
