package middleware_auth

import (
	"book-store/config"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"net/http"
)

var jwtSecret = []byte(config.GetJWTSecret())

func ValidateJWT(log *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				log.Warn("missing token")
				http.Error(w, "Missing token", http.StatusUnauthorized)
				return
			}
			if len(tokenString) < len("Bearer ") {
				log.Warn("invalid token format")
				http.Error(w, "Invalid token format", http.StatusUnauthorized)
				return
			}

			tokenString = tokenString[len("Bearer "):]
			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			})

			if err != nil || !token.Valid {
				log.Error("Invalid token", zap.String("error", err.Error()))
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			log.Info("token validation successfully")
			next.ServeHTTP(w, r)
		})
	}
}

// func Testing(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	log.Println(r.Header.Get("Authorization"))
// 	next.ServeHTTP(w, r)
// 	})
// 	}
