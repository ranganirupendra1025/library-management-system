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

    /*initailly add books and check with delete update book api ,you need to pass 
	valid username then only it will be executed.*/
	http.HandleFunc("/users", controller.GetAllUsers(DB))
	http.HandleFunc("/lms/subscriptions", controller.GetAllSubscription(DB))
	http.HandleFunc("/subscribeuser", controller.SubscribeUser(DB))
	http.HandleFunc("/registerusers", controller.RegisterUser(DB))
	http.HandleFunc("/loginuser", controller.LoginUser(DB))
	http.HandleFunc("/addbooks",controller.Addingbooks(DB))
	http.HandleFunc("/books",controller.GetAllBooks(DB))
	http.HandleFunc("/getbook/",controller.GetSingleBook(DB))
	http.HandleFunc("/update/",controller.UpdateStockBooks(DB))
	http.HandleFunc("/delete/",controller.Deletingbooks(DB))
	http.HandleFunc("/getuser/",controller.GetSingleUser(DB))
	http.HandleFunc("/deleteuser/",controller.DeletingUser(DB))

	//server setup
	err = http.ListenAndServe(":7111", nil)
	if err != nil {
		fmt.Println(err)
	}

}
