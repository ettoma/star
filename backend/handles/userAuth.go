package handles

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ettoma/star/models"
	"github.com/ettoma/star/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	newUser := &models.NewUser{}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	err := json.NewDecoder(r.Body).Decode(&newUser)
	utils.HandleWarning(err)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	utils.HandleWarning(err)

	newUser.Password = string(hashedPassword)

	fmt.Println(newUser.Password)
}

func Login(w http.ResponseWriter, r *http.Request) {
	type login struct {
		Password string `json:"password"`
		Hash     string `json:"hash"`
	}

	loginDetails := &login{}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&loginDetails)
	utils.HandleWarning(err)

	err = bcrypt.CompareHashAndPassword([]byte(loginDetails.Hash), []byte(loginDetails.Password))
	if err != nil {
		fmt.Print(err)
	}
}
