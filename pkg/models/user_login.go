package models

import (
	"fmt"
	"log"
	"golang.org/x/crypto/bcrypt"
	"mvc/pkg/types"
)

func UserLogin(email, password string, admin_id int) (string, types.ErrorMessage) {
	var errorMsg types.ErrorMessage

	db, err := Connection()
	
	if err != nil {
		errorMsg.Message = "connection error"
		log.Fatal(err)
	}
	defer db.Close()


	var hashedPassword  string
	err = db.QueryRow("SELECT hash FROM user WHERE email = ? and Admin_id =?", email,admin_id).Scan(&hashedPassword)
	if err != nil {
		errorMsg.Message = "Invalid Credentials"
		return " ", errorMsg
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {	
		errorMsg.Message = "Invalid Credentials"
		return " ", errorMsg
	}
var Name string

err = db.QueryRow("SELECT name FROM user WHERE email = ?", email).Scan(&Name)

	jwtToken, err := GenerateToken(Name)
	if err != nil {
		fmt.Println(err)
		errorMsg.Message = "token generation error"
		return "", errorMsg
		
	}
	

	errorMsg.Message = "no error"
	fmt.Println(errorMsg.Message)
	return jwtToken, errorMsg
}

