package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ettoma/star/models"
	"github.com/ettoma/star/utils"
)

func GenerateToken(username string) (token string, err error) {
	token, err = GenerateTokenString(username)
	if err != nil {
		fmt.Println(err)
	}
	return token, nil
}

func GenerateJWT(w http.ResponseWriter, r *http.Request) {
	var username map[string]string
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&username)
	utils.HandleWarning(err)

	token, err := GenerateTokenString(username["username"])
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(token)

}

func ValidateJWT(w http.ResponseWriter, r *http.Request) {
	var token map[string]string
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&token)
	utils.HandleWarning(err)
	res, err := ValidateToken(token["token"])
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.WriteJsonResponse(models.DefaultResponse{
			Message: err.Error(),
			Status:  http.StatusUnauthorized,
			Success: res,
		}, w)
	} else {
		w.WriteHeader(http.StatusAccepted)
		utils.WriteJsonResponse(models.DefaultResponse{
			Message: "Token successfully validated",
			Status:  http.StatusAccepted,
			Success: res,
		}, w)
	}

}
