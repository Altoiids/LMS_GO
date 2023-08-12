package models

import (
	"log"
	"mvc/pkg/types"
)

func AddUser(adminId int, name, email, hash, password, confirmPassword string) (string, types.ErrorMessage) {
	var errorMessage types.ErrorMessage

	db, err := Connection()
	if err != nil {
		errorMessage.Message = "connection error"
		log.Fatal(err)
	}
	defer db.Close()

	if password != confirmPassword {
		errorMessage.Message = "Passwords didn't match"
		return "", errorMessage
	}

	rows, err := db.Query("SELECT * FROM user WHERE name=? OR email=?", name, email)
	if err != nil {
		errorMessage.Message = "error"
		return "", errorMessage
	}
	defer rows.Close()

	if rows.Next() {
		errorMessage.Message = "user exists"
		return "", errorMessage
	} else {
		_, err = db.Exec("INSERT INTO user (adminId, name, email, hash) VALUES (?, ?, ?, ?)", adminId, name, email, hash)
		if err != nil {
			errorMessage.Message = "error"
			return "", errorMessage
		}
	}

	jwt, err := GenerateToken(name)
	if err != nil {
		errorMessage.Message = "token generation error"
		return "", errorMessage
	}
	errorMessage.Message = "no error"
	return jwt, errorMessage
}
