package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

// CORSMiddleware creates CORS middleware with allowed origins
func CORSMiddleware(allowedOrigins []string) func(http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:          86400, // 24 hours
	})

	return c.Handler
}
