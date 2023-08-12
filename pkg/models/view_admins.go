package models

import (
	"database/sql"
	"mvc/pkg/types"
)

func ViewAdmins(db *sql.DB) ([]types.UserRegistration, error) {
	rows, err := db.Query("SELECT name, email FROM user where adminId = 1")
	if err != nil {
		return nil, err
	}

	var admins []types.UserRegistration
	for rows.Next() {
		var admin types.UserRegistration
		err := rows.Scan(&admin.Name, &admin.Email)
		if err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}
	return admins, nil
}
