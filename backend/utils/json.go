package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJsonResponse(arg interface{}, w http.ResponseWriter) {

	json_data, err := json.Marshal(arg)
	HandleWarning(err)
	w.Write(json_data)
}
