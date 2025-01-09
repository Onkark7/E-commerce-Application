package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
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
func GetallUSer(w http.ResponseWriter, r *http.Request) {

	pageParam := r.URL.Query().Get("page")
	var page int
	page = 1
	if pageParam != "" {
		p, err := strconv.Atoi(pageParam)
		if err == nil && p > 0 {
			page = p
		}
	}

	users, err := models.GetAllUsers(page)
	if err != nil {
		http.Error(w, "Error retrieving users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"page":   page,
		"users":  users,
		"status": "success",
	})
}

func GetUserWithID(w http.ResponseWriter, r *http.Request) {
	userIdParam := r.URL.Query().Get("id")
	if userIdParam == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIdParam)
	if err != nil || userID <= 0 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := models.GetUserWithID(userID)
	if err != nil {
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "Error retrieving user", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"userID": userID,
		"user":   user,
		"status": "success",
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func UpdateUSer(w http.ResponseWriter, r *http.Request) {

}
