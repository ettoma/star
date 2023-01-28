package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

var domains = [...]string{"http://localhost:5173", "", "http://localhost:8000"}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.Println(r.Header)
		log.Printf("\n Url: %s \n Method: %s \n Content-length: %d \n\n", r.URL, r.Method, r.ContentLength)
		// log.Printf("User-agent: %s", r.UserAgent())
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

		allowedDomains := make([]string, 0)
		for _, domain := range domains {
			if origin == domain {
				allowedDomains = append(allowedDomains, origin)
				w.Header().Add("Access-Control-Allow-Origin", "*")
				w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
				w.Header().Add("Access-Control-Allow-Headers", "Authorization")
				next.ServeHTTP(w, r)
				break
			}
		}
		if len(allowedDomains) > 0 {
			fmt.Printf("Origin confirmed: %s \n\n", allowedDomains[0])
		} else if len(allowedDomains) == 0 {
			fmt.Printf("origin not allowed: %s \n\n", origin)
		}

	})
}
