package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"user/models"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserInfo

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	fmt.Println("my payload:", string(body))

	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if user.FirstName == "" {
		http.Error(w, "Missing  First Name", http.StatusBadRequest)
		return
	}

	if user.MobileNo == "" {
		http.Error(w, "Missing  Mobile Number", http.StatusBadRequest)
		return
	}

	if user.Address == "" {
		http.Error(w, "Missing Address", http.StatusBadRequest)
		return
	}

	if user.UserType == "" {
		http.Error(w, "Missing UserType", http.StatusBadRequest)
		return
	}
	if user.ClientID == 0 {
		http.Error(w, "Missing ClientID", http.StatusBadRequest)
		return
	}

	issue := models.AddData([]models.UserInfo{user})

	if issue != nil {
		http.Error(w, "Failed to add user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully",
		"userid":  fmt.Sprintf("%d", user.UserID),
	})

}
