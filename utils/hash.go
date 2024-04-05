package utils

import "golang.org/x/crypto/bcrypt"

func Hashpassword(password string)(string,error){
	bytes,err:=bcrypt.GenerateFromPassword([]byte(password),14)
	return string(bytes),err

}

func Checkpassword(password string,hashpassword string) bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hashpassword),[]byte(password))
	return err==nil

}