package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
)

// template
func createMuxMiddleware() mux.MiddlewareFunc {
	// Create a new middleware
	return func(f http.Handler) http.Handler {
		// Create a new handler encapsulate http.HandlerFunc
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Middleware Logic
			// ...
			// invoke next middleware or handler dealer
			f.ServeHTTP(w, r)
		})
	}
}
