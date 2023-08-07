package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(email, password string) (string, error) {
	db, err := Connection()
	
	if err != nil {
		return "", err
	}
	defer db.Close()


	var hashedPassword  string
	err = db.QueryRow("SELECT hash FROM user WHERE email = ? and Admin_id =1", email).Scan(&hashedPassword)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {	
		fmt.Println(err)
		return "", err
	}
var Name string
err = db.QueryRow("SELECT name FROM user WHERE email = ?", email).Scan(&Name)
	jwtToken, err := GenerateToken(Name)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	
	return jwtToken, err
}