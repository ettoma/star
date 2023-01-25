package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

var domains = [...]string{"http://localhost:5173", ""}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n Url: %s \n Method: %s \n Content-length: %d \n User-agent: %s \n", r.URL, r.Method, r.ContentLength, r.UserAgent())
		next.ServeHTTP(w, r)
	})

}

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// TODO : implement list of accepted domains

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		for _, allowedDomain := range domains {
			if origin == allowedDomain {
				w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
				fmt.Printf("origin confirmed: %s \n", r.Header.Get("Origin"))
				next.ServeHTTP(w, r)
			} else {
				fmt.Printf("origin not allowed: %s \n", origin)
			}
		}

	})
}
