package middlewares

import (
	"net/http"
	"strings"

	"github.com/ettoma/star/auth"
	"github.com/ettoma/star/models"
	"github.com/ettoma/star/utils"
)

func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			next.ServeHTTP(w, r)
		} else {
			tokenString := r.Header.Get("authorization")

			if len(tokenString) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				utils.WriteJsonResponse(models.DefaultResponse{
					Message: "Token not provided",
					Status:  http.StatusBadRequest,
					Success: false,
				}, w)
			} else {
				token := strings.Split(tokenString, " ")[1]
				res, err := auth.ValidateToken(token)
				if !res {
					w.WriteHeader(http.StatusUnauthorized)
					utils.WriteJsonResponse(models.DefaultResponse{
						Message: err.Error(),
						Success: false,
						Status:  http.StatusUnauthorized,
					}, w)
				} else {
					w.WriteHeader(http.StatusOK)
					next.ServeHTTP(w, r)
				}
			}
		}
	})
}
