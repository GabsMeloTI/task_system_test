package controllers

import (
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"awesomeProject/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var userService = service.NewUserService()

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := userService.GetUsers()
	if err != nil {
		http.Error(w, "Error retrieving users", http.StatusInternalServerError)
		return
	}

	var usersDTO []user_dto.UserListingDTO
	for _, user := range users {
		userDTO := user_dto.UserListingDTO{
			ID:           user.ID,
			Name:         user.Name,
			Email:        user.Email,
			Photo:        user.Avatar,
			CreatedAt:    user.CreatedAt,
			ProjectCount: len(user.Projects),
		}
		usersDTO = append(usersDTO, userDTO)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usersDTO)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := userService.GetUserByID(params["id"])
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	userDTO := user_dto.UserListingDTO{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Photo:        user.Avatar,
		CreatedAt:    user.CreatedAt,
		ProjectCount: len(user.Projects),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userDTO)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err = userService.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err = userService.UpdateUser(params["id"], &user)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := userService.DeleteUser(params["id"])
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err = userService.RegisterUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	token, err := userService.Login(creds.Email, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(24 * time.Hour),
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
