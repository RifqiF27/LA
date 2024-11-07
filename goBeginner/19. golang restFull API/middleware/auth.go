package middleware_auth

import (
	"book-store/config"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

// func BasicAuth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		user, pass, ok := r.BasicAuth()
// 		if !ok || user != "lumoshive" || pass != "academy" {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }

var jwtSecret = []byte(config.GetJWTSecret())

func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Authorization header:", r.Header.Get("Authorization"))
		tokenString := r.Header.Get("Authorization")
		fmt.Println("Received token:", tokenString)
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}
		if len(tokenString) < len("Bearer ") {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		tokenString = tokenString[len("Bearer "):]
		fmt.Println("Received Authorization header:", tokenString)
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
