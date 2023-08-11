package models

import (
	_ "github.com/go-sql-driver/mysql" 
	"log"

	"mvc/pkg/types"
	
	
)

func AddUser(admin_id int, name, email, hash, password, confirmPassword string ) (string, types.ErrorMessage) {
	var errorMsg types.ErrorMessage


	db, err := Connection()
	if err != nil {
		errorMsg.Message = "connection error"
		log.Fatal(err)
	}
	defer db.Close()

	if password != confirmPassword{
		errorMsg.Message = "Passwords didn't match"
		return "", errorMsg
	}
    
	

	rows, err := db.Query("SELECT * FROM user WHERE name=? OR email=?", name, email)
	if err != nil {
		errorMsg.Message = "error"
		return "", errorMsg // Return the error directly instead of a string
	}
	defer rows.Close()

	if rows.Next() {
		
		errorMsg.Message = "user exists"
		return "", errorMsg
	} else {
		_, err = db.Exec("INSERT INTO user (admin_id, name, email, hash) VALUES (?, ?, ?, ?)", admin_id, name, email, hash)
		if err != nil {
			errorMsg.Message = "error"
			return "", errorMsg // Return the error directly instead of a string
		}
	}

	

	jwt, err := GenerateToken(name)
	if err != nil {
		errorMsg.Message = "token generation error"
		return "", errorMsg// Return the error directly instead of a string
	}
	errorMsg.Message = "no error"
	return jwt, errorMsg // Return the JWT and nil error if everything is successful
}