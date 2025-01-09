package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	data "user/database"
)

type UserInfo struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	MobileNo  string    `json:"mobile_no"`
	Gender    string    `json:"gender"`
	Address   string    `json:"address"`
	UserType  string    `json:"usertype"`
	ClientID  int64     `json:"client_id"`
	Status    int32     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
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

func GetAllUsers(page int) ([]UserInfo, error) {
	var users []UserInfo

	rowsPerPage := 10
	offset := (page - 1) * rowsPerPage

	query := fmt.Sprintf("SELECT  user_id, first_name, last_name, mobile_no, gender, address, usertype, client_id, status,createdAt FROM userinfo ORDER BY createdAt DESC LIMIT %d OFFSET %d", rowsPerPage, offset)

	rows, err := data.DB.Query(query)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user UserInfo
		err := rows.Scan(
			&user.UserID,
			&user.FirstName,
			&user.LastName,
			&user.MobileNo,
			&user.Gender,
			&user.Address,
			&user.UserType,
			&user.ClientID,
			&user.Status,
			&user.CreatedAt,
		)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	}

	return users, nil
}

func GetUserWithID(userID int) ([]UserInfo, error) {
	var users []UserInfo

	query := "SELECT user_id, first_name, last_name, mobile_no, gender, address, usertype, client_id, status, createdAt FROM userinfo WHERE user_id = ?"

	rows, err := data.DB.Query(query, userID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user UserInfo

		err := rows.Scan(
			&user.UserID,
			&user.FirstName,
			&user.LastName,
			&user.MobileNo,
			&user.Gender,
			&user.Address,
			&user.UserType,
			&user.ClientID,
			&user.Status,
			&user.CreatedAt,
		)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	}

	return users, nil
}

func UpdateUser() {}
