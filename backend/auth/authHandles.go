package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func Validate(tokenToValidate string) (result bool, err error) {
	res, err := ValidateToken(tokenToValidate)

	return res, nil
}

func ValidateJWT(w http.ResponseWriter, r *http.Request) {
	var token map[string]string
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&token)
	utils.HandleWarning(err)
	res, err := ValidateToken(token["token"])
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.WriteJsonResponse(err.Error(), w)
	} else {
		w.WriteHeader(http.StatusAccepted)
		utils.WriteJsonResponse(res, w)
	}

}
