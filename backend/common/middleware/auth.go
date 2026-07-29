package middleware

import (
	"net/http"
	"strings"
)

// AuthMiddleware validates JWT tokens and adds user context to requests
// This is a universal authentication middleware that can be used across all microservices
func AuthMiddleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get token from Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)
				return
			}

			// Check if it's a Bearer token
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
				return
			}

			token := parts[1]

			// Validate token (implement your JWT validation logic here)
			// For now, we'll just check if the token is not empty
			if token == "" {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// Add user context to request (implement your context logic here)
			// ctx := context.WithValue(r.Context(), "user", userClaims)
			// r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// OptionalAuthMiddleware validates JWT tokens if present, but doesn't require them
// Useful for endpoints that work both with and without authentication
func OptionalAuthMiddleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader != "" {
				parts := strings.Split(authHeader, " ")
				if len(parts) == 2 && parts[0] == "Bearer" {
					token := parts[1]
					// Validate token and add user context if valid
					// ctx := context.WithValue(r.Context(), "user", userClaims)
					// r = r.WithContext(ctx)
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}

// RoleMiddleware checks if the user has the required role
// This should be used after AuthMiddleware
func RoleMiddleware(requiredRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get user from context (set by AuthMiddleware)
			// user := r.Context().Value("user")

			// Check if user has required role
			// Implement your role checking logic here

			// For now, we'll just pass through
			next.ServeHTTP(w, r)
		})
	}
}
