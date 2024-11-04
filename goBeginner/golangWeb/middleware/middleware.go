package middleware

import (
	"encoding/json"
	"net/http"
	"todo/model"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("token")
		if authHeader != "12345" {
			w.WriteHeader(http.StatusUnauthorized)
			badResponse := model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized",
				Data:       nil,
			}
			json.NewEncoder(w).Encode(badResponse)
			return
		}

		next.ServeHTTP(w, r)
	})
}
