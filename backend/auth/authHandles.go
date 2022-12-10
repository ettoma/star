package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ettoma/star/utils"
)

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
		fmt.Println(err)
	}
	utils.WriteJsonResponse(res, w)

}