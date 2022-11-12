package handles

import (
	"encoding/json"
	"net/http"

	"github.com/ettoma/star/database"
	"github.com/ettoma/star/models"
	"github.com/ettoma/star/utils"
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
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusCreated)
		data, err := json.Marshal(createdUser)
		utils.HandleWarning(err)
		w.Write(data)

	}

}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	var id map[string]int

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&id)
	utils.HandleFatal(err)

	database.DeleteUserById(id["id"])
}

func DeleteUserByUsername(w http.ResponseWriter, r *http.Request) {
	var username map[string]string

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&username)
	utils.HandleFatal(err)

	database.DeleteUserByUsername(username["username"])
}
