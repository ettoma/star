package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ettoma/star/models"
)

func HandleWarning(err error) {
	if err != nil {
		log.Println(err)
	}
}

func HandleFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func HandleNotFound(w http.ResponseWriter, searchKey string) {

	response := models.DefaultResponse{
		Message: fmt.Sprintf("%s not found", searchKey),
		Status:  http.StatusNotFound,
		Success: false,
	}
	responseJson, err := json.Marshal(response)
	HandleWarning(err)
	w.Write(responseJson)
}
