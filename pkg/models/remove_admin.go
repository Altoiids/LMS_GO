package models

import (
	"log"
)


func RemoveAdmin( email string) (string){
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("DELETE from user WHERE email = ?", email)
	if err != nil {
		return "There was error"
	}
	defer rows.Close()
	
	return ""
}