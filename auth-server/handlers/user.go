package handlers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ishtiaqhimel/oms/auth-server/initializers"
	"github.com/ishtiaqhimel/oms/auth-server/models"
	"golang.org/x/crypto/bcrypt"
)

// swagger:route POST /signup signupUser
// SignUp a new user
//
// responses:
//		201: noContentResponse
//	 400: errorResponse
//	 500: errorResponse

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	log.Println("Creating a new user")

	user := models.User{}
	if err := models.FromJSON(&user, r.Body); err != nil {
		log.Println("Error deserializing", err)
		w.WriteHeader(http.StatusBadRequest)
		models.ToJSON(&models.GenericError{Message: err.Error()}, w)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error while hashing password", err)
		w.WriteHeader(http.StatusInternalServerError)
		models.ToJSON(&models.GenericError{Message: err.Error()}, w)
		return
	}

	user.Password = string(hashedPassword)
	if err := initializers.DB.Create(&user).Error; err != nil {
		log.Println("Error while creating user", err)
		w.WriteHeader(http.StatusInternalServerError)
		models.ToJSON(&models.GenericError{Message: err.Error()}, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// swagger:route POST /login loginUser
// Login an user
//
// responses:
//		200: noContentResponse
//	 400: errorResponse
//	 401: errorResponse
//	 500: errorResponse

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	user := models.User{}
	if err := models.FromJSON(&user, r.Body); err != nil {
		log.Println("Error deserializing", err)
		w.WriteHeader(http.StatusBadRequest)
		models.ToJSON(&models.GenericError{Message: err.Error()}, w)
		return
	}

	username, password := user.Username, user.Password

	err := initializers.DB.First(&user, "username = ?", username).Error
	if err != nil {
		log.Println("Invalid credentials", err)
		w.WriteHeader(http.StatusUnauthorized)
		models.ToJSON(&models.GenericError{Message: err.Error()}, w)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println("Invalid credentials", err)
		w.WriteHeader(http.StatusUnauthorized)
		models.ToJSON(&models.GenericError{Message: err.Error()}, w)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": os.Getenv("JWT_KEY"),
		"sub": user.Username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println("Error while generating token", err)
		w.WriteHeader(http.StatusInternalServerError)
		models.ToJSON(&models.GenericError{Message: err.Error()}, w)
		return
	}

	w.Header().Add("Authorization", "Bearer "+tokenString)
}
