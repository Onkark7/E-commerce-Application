package models

import (
	"fmt"
	"log"
	data "user/database"
)

type Usertype struct {
	Usertype_id int    `json:"usertype_id" `
	Type        string `json:"type"`
}

func Getusertype() ([]Usertype, error) {
	var userTypes []Usertype
	getq := fmt.Sprintf("select usertype_id,type FROM usertype")

	rows, err := data.DB.Query(getq)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var Usertype Usertype
		err := rows.Scan(
			&Usertype.Usertype_id,
			&Usertype.Type,
		)

		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		userTypes = append(userTypes, Usertype)
	}

	return userTypes, nil
}
