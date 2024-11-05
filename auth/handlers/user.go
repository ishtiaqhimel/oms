package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ishtiaqhimel/oms/auth/initializers"
	"github.com/ishtiaqhimel/oms/auth/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Error while decode request body", http.StatusInternalServerError)
		return
	}

	username := data["username"]
	password := data["password"]

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error while hashing password", http.StatusInternalServerError)
		return
	}

	user := models.User{Username: username, Password: string(hashedPassword)}
	if err := initializers.DB.Create(&user).Error; err != nil {
		http.Error(w, "Error while creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Error while decode request body", http.StatusInternalServerError)
		return
	}

	username := data["username"]
	password := data["password"]

	var user models.User
	err := initializers.DB.First(&user, "username = ?", username).Error
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error while generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Add("token", tokenString)
	w.Write([]byte("Logged in successfully"))
}
