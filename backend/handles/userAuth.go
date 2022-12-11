package handles

import (
	"encoding/json"
	"net/http"

	"github.com/ettoma/star/database"
	"github.com/ettoma/star/models"
	"github.com/ettoma/star/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	newUser := &models.NewUser{}
	success := false

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	err := json.NewDecoder(r.Body).Decode(&newUser)
	utils.HandleWarning(err)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	utils.HandleWarning(err)

	newUser.Password = string(hashedPassword)

	createdUser, err := database.AddUser(newUser.Name, newUser.Username)
	if err != nil {
		if err.Error() == "name or username too short (min. 4 char)" {
			w.WriteHeader(http.StatusBadRequest)
			response := models.DefaultResponse{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
				Success: false,
			}
			utils.WriteJsonResponse(response, w)
		} else {
			w.WriteHeader(http.StatusConflict)
			response := models.DefaultResponse{
				Message: err.Error(),
				Status:  http.StatusConflict,
				Success: false,
			}
			utils.WriteJsonResponse(response, w)
		}
	} else {
		success = true
		w.WriteHeader(http.StatusCreated)
		utils.WriteJsonResponse(createdUser, w)
	}

	if success {
		err := database.AddUserToAuth(newUser.Password)
		utils.HandleWarning(err)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	loginDetails := &models.Login{}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&loginDetails)
	utils.HandleWarning(err)

	hashedPassword, err := database.GetHashForUser(loginDetails.Username)
	if err != nil {
		utils.HandleNotFound(w, loginDetails.Username)
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(loginDetails.Password))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := models.DefaultResponse{
				Message: "Incorrect password",
				Status:  http.StatusBadRequest,
				Success: false,
			}
			utils.WriteJsonResponse(response, w)
		} else {
			w.WriteHeader(http.StatusOK)
			response := models.DefaultResponse{
				Message: "authentication successful",
				Status:  http.StatusOK,
				Success: true,
			}
			utils.WriteJsonResponse(response, w)
		}
	}

}
