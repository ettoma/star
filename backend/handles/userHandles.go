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

func AddUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var newUser models.User

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&newUser)
	utils.HandleWarning(err)

	createdUser, err := database.AddUser(newUser.Name, newUser.Username)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		response := models.DefaultResponse{
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

	id := mux.Vars(r)

	idInt, err := strconv.Atoi(id["id"])
	utils.HandleWarning(err)

	user, err := database.GetUserById(idInt)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := models.DefaultResponse{
			Message: "User not found",
			Status:  http.StatusNotFound,
			Success: false,
		}
		responseJson, err := json.Marshal(response)
		utils.HandleWarning(err)
		w.Write(responseJson)
	} else {
		w.WriteHeader(http.StatusOK)
		responseUser := models.User{
			Name:      user.Name,
			Username:  user.Username,
			Id:        user.Id,
			CreatedAt: user.CreatedAt,
		}
		userJson, err := json.Marshal(responseUser)
		utils.HandleWarning(err)
		w.Write(userJson)
	}
}

// TODO implement this:
func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	// vars := mux.Vars(r)

	// qUsername := vars["username"]
	// // utils.HandleWarning(err)

	// user, err := database.GetUserByUsername(qUsername)

	// if err != nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	response := models.DefaultResponse{
	// 		Message: "User not found",
	// 		Status:  http.StatusNotFound,
	// 		Success: false,
	// 	}
	// 	responseJson, err := json.Marshal(response)
	// 	utils.HandleWarning(err)
	// 	w.Write(responseJson)
	// } else {
	// 	w.WriteHeader(http.StatusOK)
	// 	responseUser := models.User{
	// 		Name:      user.Name,
	// 		Username:  user.Username,
	// 		Id:        user.Id,
	// 		CreatedAt: user.CreatedAt,
	// 	}
	// 	userJson, err := json.Marshal(responseUser)
	// 	utils.HandleWarning(err)
	// 	w.Write(userJson)
	// }
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
		response := models.DefaultResponse{
			Message: fmt.Sprintf("user with id #%d deleted", id["id"]),
			Status:  http.StatusOK,
			Success: success,
		}
		responseJson, err := json.Marshal(response)
		utils.HandleWarning(err)
		w.Write(responseJson)
	} else {
		w.WriteHeader(http.StatusNotFound)
		response := models.DefaultResponse{
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
		response := models.DefaultResponse{
			Message: fmt.Sprintf("user with username: %s deleted", username["username"]),
			Status:  http.StatusOK,
			Success: success,
		}
		responseJson, err := json.Marshal(response)
		utils.HandleWarning(err)
		w.Write(responseJson)
	} else {
		w.WriteHeader(http.StatusNotFound)
		response := models.DefaultResponse{
			Message: fmt.Sprintf("user with username: %s not found", username["username"]),
			Status:  http.StatusNotFound,
			Success: success,
		}
		responseJson, err := json.Marshal(response)
		utils.HandleWarning(err)
		w.Write(responseJson)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

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
			responseJson, err := json.Marshal(response)
			utils.HandleWarning(err)
			w.Write(responseJson)
			break
		}
	}

}

func handleSuccess(success bool, w http.ResponseWriter) {
	if success {
		w.WriteHeader(http.StatusOK)
		response := models.DefaultResponse{
			Message: "user deleted",
			Status:  http.StatusOK,
			Success: success,
		}
		responseJson, err := json.Marshal(response)
		utils.HandleWarning(err)
		w.Write(responseJson)
	} else {
		w.WriteHeader(http.StatusNotFound)
		response := models.DefaultResponse{
			Message: "user not found",
			Status:  http.StatusNotFound,
			Success: success,
		}
		responseJson, err := json.Marshal(response)
		utils.HandleWarning(err)
		w.Write(responseJson)
	}
}
