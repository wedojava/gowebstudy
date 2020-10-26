package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Logging() mux.MiddlewareFunc {
	// Create middleware
	return func(next http.Handler) http.Handler {
		// Create a new handler encapsulate http.HandlerFunc
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				start := time.Now()
				defer func() { log.Println(r.URL.Path, time.Since(start)) }()
				next.ServeHTTP(w, r)
			})
	}

}
