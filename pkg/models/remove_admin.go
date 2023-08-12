package models

func RemoveAdmin( email string) (error){
	db, err := Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("DELETE from user WHERE email = ?", email)
	if err != nil {
		return err
	}
	defer rows.Close()
	
	return nil
}