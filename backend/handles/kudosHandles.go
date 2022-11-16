package handles

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ettoma/star/database"
	"github.com/ettoma/star/models"
	"github.com/ettoma/star/utils"
)

func AddKudos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var newKudos models.Kudos

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	// TODO : fix json unmarshaling
	err := json.NewDecoder(r.Body).Decode(&newKudos)
	utils.HandleWarning(err)

	sender, err := database.GetUserByUsername(newKudos.Sender.Username)
	receiver, err := database.GetUserByUsername(newKudos.Receiver.Username)

	fmt.Print(sender, receiver)

	kudos, err := database.AddKudos(sender, receiver, newKudos.Content)

	fmt.Print(kudos)
}
