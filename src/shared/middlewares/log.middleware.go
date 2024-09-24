package middleware

import (
	"app/src/lib/logger"
	"net/http"
	"time"
)

var log = logger.NewLogger("LoggingMiddleware")

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("Received request: %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf("Handled request: %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}
