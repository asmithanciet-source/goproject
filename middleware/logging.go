package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging logs the method, path and processing time of every request.
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started  %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf("Finished %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}
