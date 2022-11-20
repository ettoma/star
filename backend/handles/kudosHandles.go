package handles

import (
	"encoding/json"
	"net/http"

	"github.com/ettoma/star/database"
	"github.com/ettoma/star/models"
	"github.com/ettoma/star/utils"
)

func AddKudos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var newKudos models.Kudos

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&newKudos)
	utils.HandleWarning(err)

	sender, _ := database.GetUserByUsername(newKudos.Sender)
	if sender == nil {
		utils.HandleUserNotFound(w, "Sender")
	}
	receiver, _ := database.GetUserByUsername(newKudos.Receiver)
	if receiver == nil {
		utils.HandleUserNotFound(w, "Receiver")
	}

	if sender != nil && receiver != nil {
		kudos, err := database.AddKudos(newKudos.Sender, newKudos.Receiver, newKudos.Content)
		utils.HandleWarning(err)
		utils.WriteJsonResponse(kudos, w)
	}

}

func GetAllKudos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kudos := database.GetAllKudos()

	utils.WriteJsonResponse(kudos, w)

}

func GetKudosPerUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user map[string]string
	var kudos []*models.Kudos

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&user)
	utils.HandleWarning(err)

	if user["sender"] != "" {

		kudos, err = database.GetKudosPerSender(user["sender"])
		if err != nil {
			utils.HandleUserNotFound(w, "Sender")
		} else {
			utils.WriteJsonResponse(kudos, w)
		}

	} else if user["receiver"] != "" {

		kudos, err = database.GetKudosPerReceiver(user["receiver"])
		if err != nil {
			utils.HandleUserNotFound(w, "Receiver")
		} else {
			utils.WriteJsonResponse(kudos, w)
		}
	}

}
