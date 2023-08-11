package models

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"mvc/pkg/types"
)

func ViewAdmins(db *sql.DB) ([]types.UserReg, error) {
	
	rows, err := db.Query("SELECT name, email FROM user where Admin_id = 1")
	if err != nil {
		return nil, err
	}

	var admins []types.UserReg
	for rows.Next() {
		var admin types.UserReg
		err := rows.Scan(&admin.Name, &admin.Email)
		if err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}
	return admins, nil
}