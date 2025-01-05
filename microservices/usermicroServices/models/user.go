package models

import (
	"encoding/json"
	"fmt"
	"log"
	data "user/database"
)

type UserInfo struct {
	id        int    `json:"id"`
	UserID    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	MobileNo  string `json:"mobile_no"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	UserType  string `json:"usertype"`
	ClientID  int64  `json:"client_id"`
	Status    int32  `json:"status"`
}

func AddData(userinfor []UserInfo) error {
	User_information, err := json.Marshal(userinfor)

	if err != nil {
		log.Fatalf("userinformation not send properly %v", err)
	}

	Query := "insert into userinfo (user_id,first_name,last_name,mobile_no,gender,address,usertype,client_id,status) values (?,?,?,?,?,?,?,?,?)"

	for _, user := range userinfor {

		_, err := data.DB.Exec(Query, user.UserID, user.FirstName, user.LastName, user.MobileNo, user.Gender, user.Address, user.UserType, user.ClientID, user.Status)

		if err != nil {
			log.Fatalf("error in adding Userinformation in datbase %v", err)
			log.Printf("Error executing query: %s with data %+v. Error: %v", Query, user, err)
			return fmt.Errorf("error inserting user data: %w", err)

		} else {
			fmt.Printf("Successfully inserted user with user_id %d%s\n", user.UserID, User_information)
		}
	}
	return nil
}
