package handles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ettoma/star/auth"
	"github.com/ettoma/star/database"
	"github.com/ettoma/star/models"
	"github.com/ettoma/star/utils"
	"github.com/gorilla/mux"
)

func AddKudos(w http.ResponseWriter, r *http.Request) {

	var newKudos models.Kudos

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&newKudos)
	utils.HandleWarning(err)

	sender, _ := database.GetUserByUsername(newKudos.Sender)
	if sender == nil {
		utils.HandleNotFound(w, "Sender")
	}
	receiver, _ := database.GetUserByUsername(newKudos.Receiver)
	if receiver == nil {
		utils.HandleNotFound(w, "Receiver")
	}

	if len(newKudos.Content) == 0 {
		utils.HandleNotFound(w, "Content")
	}

	if sender != nil && receiver != nil && len(newKudos.Content) > 0 {
		kudos, err := database.AddKudos(newKudos.Sender, newKudos.Receiver, newKudos.Content)
		utils.HandleWarning(err)
		utils.WriteJsonResponse(kudos, w)
	}

}

func GetAllKudos(w http.ResponseWriter, r *http.Request) {

	kudos := database.GetAllKudos()

	utils.WriteJsonResponse(kudos, w)

}

func GetKudosPerUser(w http.ResponseWriter, r *http.Request) {

	var receiver = mux.Vars(r)["receiver"]
	var tokenString = r.Header.Get("authorization")
	// fmt.Println("my token string: ", tokenString)
	if len(tokenString) == 0 {
		tokenString, _ = auth.GenerateTokenString(receiver)
	} else {
		tokenString = strings.Split(tokenString, " ")[1]
	}

	var kudos []*models.Kudos
	// TODO! implement user check before response is sent, make sure user can only access his own data

	_, err := auth.UserInQueryMatchToken(tokenString, receiver)

	if err != nil {
		utils.WriteJsonResponse(models.DefaultResponse{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
			Success: false,
		}, w)
	} else {
		kudos, err = database.GetKudosPerReceiver(receiver)
		if err != nil {
			utils.HandleNotFound(w, "Receiver")
		} else {
			utils.WriteJsonResponse(kudos, w)
		}

	}

}

func DeleteKudos(w http.ResponseWriter, r *http.Request) {
	var id map[string]int

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(&id)
	utils.HandleWarning(err)

	success, err := database.DeleteKudos(id["id"])

	if success {
		w.WriteHeader(http.StatusOK)
		response := models.DefaultResponse{
			Message: "kudos deleted",
			Status:  http.StatusOK,
			Success: success,
		}
		utils.WriteJsonResponse(response, w)
	} else {
		utils.HandleNotFound(w, fmt.Sprintf("Kudos with id %d", id["id"]))
	}

}
