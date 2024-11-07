package middleware

import (
	"book-store/config"
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
        tokenString := r.Header.Get("Authorization")
        if tokenString == "" {
            http.Error(w, "Missing token", http.StatusUnauthorized)
            return
        }

        // Ambil token tanpa prefix "Bearer "
        tokenString = tokenString[len("Bearer "):]

        // Verifikasi token
        claims := jwt.MapClaims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Token valid, lanjutkan ke handler berikutnya
        next.ServeHTTP(w, r)
    })
}
