package main

import (
	"io/ioutil"
	"lms/db"
	"log"
)

func main(){
	//connecting to databse
	DB,err:=db.InitDB()
	if err!=nil{
		log.Fatal("error in connecting")
	}
	defer DB.Close()
	script,err:=ioutil.ReadFile("db/init.sql")
	if err!=nil{
		//fmt.Println(err)
		log.Fatal(err)
	}
	_,err=DB.Exec(string(script))
	if err!=nil{
		log.Fatal(err)
	}
	

}