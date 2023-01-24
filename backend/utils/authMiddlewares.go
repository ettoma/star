package utils

import (
	"log"
	"net/http"
)

func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("validating token")
		next.ServeHTTP(w, r)
	})
}
