package handles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ettoma/star/database"
	"github.com/ettoma/star/models"
	"github.com/ettoma/star/utils"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users := database.GetAllUsers()

	utils.WriteJsonResponse(users, w)
}

func AddUser(w http.ResponseWriter, r *http.Request) {

	var newUser models.User

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&newUser)
	utils.HandleWarning(err)

	createdUser, err := database.AddUser(newUser.Name, newUser.Username, "")
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
		w.WriteHeader(http.StatusCreated)
		utils.WriteJsonResponse(createdUser, w)

	}

}

func GetUserById(w http.ResponseWriter, r *http.Request) {

	var id map[string]int

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&id)
	utils.HandleFatal(err)

	user, err := database.GetUserById(id["id"])

	if err != nil {
		utils.HandleNotFound(w, fmt.Sprintf("User with id %d", id["id"]))
	} else {
		w.WriteHeader(http.StatusOK)
		responseUser := models.User{
			Name:      user.Name,
			Username:  user.Username,
			Id:        user.Id,
			CreatedAt: user.CreatedAt,
		}
		utils.WriteJsonResponse(responseUser, w)
	}
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {

	var id map[string]int

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&id)
	utils.HandleFatal(err)

	success, err := database.DeleteUserById(id["id"])

	if success {
		w.WriteHeader(http.StatusOK)
		response := models.DefaultResponse{
			Message: fmt.Sprintf("user with id #%d deleted", id["id"]),
			Status:  http.StatusOK,
			Success: success,
		}
		utils.WriteJsonResponse(response, w)
	} else {
		utils.HandleNotFound(w, fmt.Sprintf("User with id %d", id["id"]))
	}

}

func DeleteUserByUsername(w http.ResponseWriter, r *http.Request) {

	var username map[string]string

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&username)
	utils.HandleFatal(err)

	success, err := database.DeleteUserByUsername(username["username"])

	if success {
		w.WriteHeader(http.StatusOK)
		response := models.DefaultResponse{
			Message: fmt.Sprintf("user with username: %s deleted", username["username"]),
			Status:  http.StatusOK,
			Success: success,
		}
		utils.WriteJsonResponse(response, w)
	} else {
		utils.HandleNotFound(w, fmt.Sprintf("User with username %s", username["username"]))
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	var u map[string]interface{}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&u)
	utils.HandleFatal(err)

	var success bool

	for k := range u {
		if k == "username" {
			success, err = database.DeleteUserByUsername(fmt.Sprintf("%s", u["username"]))
			handleSuccess(success, w)
			break
		} else if k == "id" {
			value, err := strconv.Atoi(fmt.Sprintf("%v", u["id"]))
			utils.HandleWarning(err)
			success, err = database.DeleteUserById(value)
			handleSuccess(success, w)
			break
		} else {
			w.WriteHeader(http.StatusBadRequest)
			response := models.DefaultResponse{
				Message: "invalid or malformed request",
				Status:  http.StatusBadRequest,
				Success: success,
			}
			utils.WriteJsonResponse(response, w)
			break
		}
	}

}

// func PatchUser(w http.ResponseWriter, r *http.Request) {
// 	var req map[string]interface{}

// 	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	utils.HandleWarning(err)

// 	var success bool

// 	if req["id"] > 0 {

// 	}
// }

func handleSuccess(success bool, w http.ResponseWriter) {
	if success {
		w.WriteHeader(http.StatusOK)
		response := models.DefaultResponse{
			Message: "user deleted",
			Status:  http.StatusOK,
			Success: success,
		}
		utils.WriteJsonResponse(response, w)
	} else {
		w.WriteHeader(http.StatusNotFound)
		response := models.DefaultResponse{
			Message: "user not found",
			Status:  http.StatusNotFound,
			Success: success,
		}
		utils.WriteJsonResponse(response, w)
	}
}
