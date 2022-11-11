package handles

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ettoma/star/database"
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
	req := r.Body

	reqBody, err := io.ReadAll(req)
	utils.HandleWarning(err)

	var rtest string

	json.Unmarshal(reqBody, &rtest)

	fmt.Print(rtest)

}
