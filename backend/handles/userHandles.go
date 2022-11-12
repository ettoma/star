package handles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ettoma/star/database"
	"github.com/ettoma/star/models"
	"github.com/ettoma/star/utils"
	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users := database.GetAllUsers()

	w.Header().Set("Content-Type", "application/json")

	json_data, err := json.Marshal(users)
	utils.HandleWarning(err)
	w.Write(json_data)
}

// TODO : implement timestamp for creation
func AddUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var newUser models.User

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&newUser)
	utils.HandleWarning(err)

	createdUser, err := database.AddUser(newUser.Name, newUser.Username)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		response := models.SimpleResponse{
			Message: err.Error(),
			Status:  http.StatusConflict,
			Success: false,
		}
		responseJson, err := json.Marshal(response)
		utils.HandleWarning(err)
		w.Write(responseJson)
	} else {
		w.WriteHeader(http.StatusCreated)
		data, err := json.Marshal(createdUser)
		utils.HandleWarning(err)
		w.Write(data)

	}

}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users = database.GetAllUsers()

	id := mux.Vars(r)

	idInt, err := strconv.Atoi(id["id"])
	utils.HandleWarning(err)

	for _, user := range users {
		if user.Id == idInt {
			w.WriteHeader(http.StatusOK)
			responseUser := models.User{
				Name:     user.Name,
				Username: user.Username,
				Id:       user.Id,
			}
			userJson, err := json.Marshal(responseUser)
			utils.HandleWarning(err)
			w.Write(userJson)
		} else {
			w.WriteHeader(http.StatusNotFound)
			response := models.SimpleResponse{
				Message: "User not found",
				Status:  http.StatusNotFound,
				Success: false,
			}
			responseJson, err := json.Marshal(response)
			utils.HandleWarning(err)
			w.Write(responseJson)
		}
	}

}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var id map[string]int

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&id)
	utils.HandleFatal(err)

	success, err := database.DeleteUserById(id["id"])

	if success {
		w.WriteHeader(http.StatusOK)
		response := models.SimpleResponse{
			Message: fmt.Sprintf("user with id #%d deleted", id["id"]),
			Status:  http.StatusOK,
			Success: success,
		}
		responseJson, err := json.Marshal(response)
		utils.HandleWarning(err)
		w.Write(responseJson)
	} else {
		w.WriteHeader(http.StatusNotFound)
		response := models.SimpleResponse{
			Message: fmt.Sprintf("user with id #%d not found", id["id"]),
			Status:  http.StatusNotFound,
			Success: success,
		}
		responseJson, err := json.Marshal(response)
		utils.HandleWarning(err)
		w.Write(responseJson)
	}

}

func DeleteUserByUsername(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var username map[string]string

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&username)
	utils.HandleFatal(err)

	success, err := database.DeleteUserByUsername(username["username"])

	if success {
		w.WriteHeader(http.StatusOK)
		response := models.SimpleResponse{
			Message: fmt.Sprintf("user with username: %s deleted", username["username"]),
			Status:  http.StatusOK,
			Success: success,
		}
		responseJson, err := json.Marshal(response)
		utils.HandleWarning(err)
		w.Write(responseJson)
	} else {
		w.WriteHeader(http.StatusNotFound)
		response := models.SimpleResponse{
			Message: fmt.Sprintf("user with username: %s not found", username["username"]),
			Status:  http.StatusNotFound,
			Success: success,
		}
		responseJson, err := json.Marshal(response)
		utils.HandleWarning(err)
		w.Write(responseJson)
	}
}
