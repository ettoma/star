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

	err := json.NewDecoder(r.Body).Decode(&newKudos)
	utils.HandleWarning(err)

	kudos, err := database.AddKudos(newKudos.Sender, newKudos.Receiver, newKudos.Content)

	fmt.Print(kudos)
}

func GetAllKudos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kudos := database.GetAllKudos()

	json_data, err := json.Marshal(kudos)
	utils.HandleWarning(err)
	w.Write(json_data)

}
