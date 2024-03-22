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

	//this was sample api for getting users from databse using this u can change as per ur need
	http.HandleFunc("/users", controller.GetAllUsers(DB))
	http.HandleFunc("/registerusers", controller.RegisterUser(DB))
	//server setup
	err = http.ListenAndServe(":7111", nil)
	if err != nil {
		fmt.Println(err)
	}

}
