package models

import (
	"log"
	"mvc/pkg/types"
	"golang.org/x/crypto/bcrypt"
)

func UserLogin(email, password string, adminId int) (string, types.ErrorMessage) {
	var errorMessage types.ErrorMessage

	db, err := Connection()
	if err != nil {
		errorMessage.Message = "connection error"
		log.Fatal(err)
	}
	defer db.Close()

	var hashedPassword string
	err = db.QueryRow("SELECT hash FROM user WHERE email = ? and adminId =?", email, adminId).Scan(&hashedPassword)
	if err != nil {
		errorMessage.Message = "Invalid Credentials"
		return " ", errorMessage
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		errorMessage.Message = "Invalid Credentials"
		return " ", errorMessage
	}

	var Name string
	err = db.QueryRow("SELECT name FROM user WHERE email = ?", email).Scan(&Name)
	jwtToken, err := GenerateToken(Name)
	if err != nil {
		errorMessage.Message = "Token generation error"
		return "", errorMessage
	}

	errorMessage.Message = "no error"
	return jwtToken, errorMessage
}
