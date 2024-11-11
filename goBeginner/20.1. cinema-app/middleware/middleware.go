package middleware

import (
	"cinema/service"
	"cinema/utils"
	"net/http"
)

// AuthMiddleware checks if the token is valid and is used as a middleware
func AuthMiddleware(authService service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the token from the Authorization header
			token := r.Header.Get("Authorization")
			if len(token) > 7 && token[:7] == "Bearer " {
				token = token[7:] // Remove the "Bearer " prefix
			}

			// Check if the token is empty
			if token == "" {
				utils.SendJSONResponse(w, http.StatusUnauthorized, "Token required", nil)
				return
			}

			// Verify the token
			_, err := authService.VerifyToken(token)
			if err != nil {
				utils.SendJSONResponse(w, http.StatusUnauthorized, "Invalid or expired token", nil)
				return
			}

			// Proceed with the next handler if token is valid
			next.ServeHTTP(w, r)
		})
	}
}
