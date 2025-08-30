package middlewares

import (
	"backend/utils"
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// ContextKey é um tipo personalizado para evitar colisões no context
type ContextKey string

const UserIDKey ContextKey = "user_id"

// AuthenticationMiddleware checks if the user has a valid JWT token
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Missing authentication token"})
			return
		}

		// The token should be prefixed with "Bearer "
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid authentication token"})
			return
		}

		tokenString = tokenParts[1]

		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid authentication token"})
			return
		}

		// Add user_id to request context
		ctx := context.WithValue(r.Context(), UserIDKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserIDFromContext extrai o user_id do context da requisição
func GetUserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(UserIDKey).(string)
	return userID, ok
}
