package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/ishtiaqhimel/oms/oms-server/models"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Token")
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				log.Println("Unauthorized", err)
				w.WriteHeader(http.StatusUnauthorized)
				models.ToJSON(&models.GenericError{Message: err.Error()}, w)
				return
			}
			log.Println("Bad Request", err)
			w.WriteHeader(http.StatusBadRequest)
			models.ToJSON(&models.GenericError{Message: err.Error()}, w)
			return
		}

		if !token.Valid {
			log.Println("Unauthorized", err)
			w.WriteHeader(http.StatusUnauthorized)
			models.ToJSON(&models.GenericError{Message: "token is not valid"}, w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
