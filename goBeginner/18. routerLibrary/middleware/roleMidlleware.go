package middleware_test


import (
	"encoding/json"
	"net/http"
	"todo/model"
)

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("role")

		if authHeader != "admin_token" {
			w.WriteHeader(http.StatusUnauthorized)
			badResponse := model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized - Admin access required",
				Data:       nil,
			}
			json.NewEncoder(w).Encode(badResponse)
			return
		}

		next.ServeHTTP(w, r)
	})
}
