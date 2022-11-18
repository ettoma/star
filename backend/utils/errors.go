package utils

import (
	"encoding/json"
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

func HandleUserNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	response := models.DefaultResponse{
		Message: "User not found",
		Status:  http.StatusNotFound,
		Success: false,
	}
	responseJson, err := json.Marshal(response)
	HandleWarning(err)
	w.Write(responseJson)
}
