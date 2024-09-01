package controllers

import (
	"awesomeProject/db"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"awesomeProject/utils"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var jwtKey = []byte("your_secret_key")

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Preload("Projects").Find(&users)

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
	var user models.User
	err := db.DB.Preload("Projects").First(&user, params["id"]).Error
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

	var existingUser models.User
	if err := db.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		http.Error(w, "Email already registered", http.StatusConflict)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error encrypting password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	err = db.DB.Create(&user).Error
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User

	err := db.DB.First(&user, params["id"]).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	db.DB.Save(&user)

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := db.DB.Delete(&models.User{}, params["id"]).Error
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

	var existingUser models.User
	if err := db.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		http.Error(w, "Email already registered", http.StatusConflict)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error encrypting password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	err = db.DB.Create(&user).Error
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
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

	var user models.User
	err = db.DB.Where("email = ?", creds.Email).First(&user).Error
	if err != nil || !utils.ComparePasswords(user.Password, creds.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &utils.JWTClaims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
