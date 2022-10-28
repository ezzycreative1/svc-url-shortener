package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func LoadConfig(mux http.Handler) *http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodOptions, http.MethodPatch},
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           86000,
		// Enable Debugging for testing, consider disabling in production
		//Debug: true,
	})
	handler := c.Handler(mux)
	return &handler
}
