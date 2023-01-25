package middlewares

import (
	"net/http"
	"strings"

	"github.com/ettoma/star/auth"
	"github.com/ettoma/star/utils"
)

func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := strings.Split(r.Header.Get("Authorization"), " ")[1]

		res, err := auth.ValidateToken(token)
		if !res {
			w.WriteHeader(http.StatusUnauthorized)
			utils.WriteJsonResponse(err.Error(), w)
		} else {
			w.WriteHeader(http.StatusAccepted)
			// utils.WriteJsonResponse(res, w)
			next.ServeHTTP(w, r)
		}
	})
}
