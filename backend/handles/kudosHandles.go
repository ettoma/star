package handles

import (
	"encoding/json"
	"net/http"

	"github.com/ettoma/star/database"
	"github.com/ettoma/star/models"
	"github.com/ettoma/star/utils"
	"github.com/gorilla/mux"
)

func AddKudos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var newKudos models.Kudos

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&newKudos)
	utils.HandleWarning(err)

	_, err = database.GetUserByUsername(newKudos.Sender)
	if err != nil {
		utils.HandleUserNotFound(w)

	}
	_, err = database.GetUserByUsername(newKudos.Receiver)
	if err != nil {
		utils.HandleUserNotFound(w)
	} else {

		kudos, err := database.AddKudos(newKudos.Sender, newKudos.Receiver, newKudos.Content)

		json_data, err := json.Marshal(kudos)
		utils.HandleWarning(err)
		w.Write(json_data)
	}

}

func GetAllKudos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kudos := database.GetAllKudos()

	json_data, err := json.Marshal(kudos)
	utils.HandleWarning(err)
	w.Write(json_data)

}

func GetKudosPerUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userMap := mux.Vars(r)
	user := userMap["username"]

	kudos, err := database.GetKudosPerUser(user)

	if err != nil {
		utils.HandleUserNotFound(w)
	} else {

		json_data, err := json.Marshal(kudos)
		utils.HandleWarning(err)
		w.Write(json_data)
	}

}
