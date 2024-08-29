package core

import (
	"encoding/json"
	"net/http"
)

func convertJSONData(data interface{}) (error, []byte) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err, []byte(err.Error())
	}
	return nil, jsonData
}

func renderJSON(w http.ResponseWriter, data interface{}, status int) {
	err, jsonData := convertJSONData(data)
	w.Header().Set("Content-Type", "application-json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(status)
	}
	w.Write(jsonData)
}

func HttpResponseSuccess(w http.ResponseWriter, data interface{}) {
	renderJSON(w, data, http.StatusOK)
}

func HttpResponseCreated(w http.ResponseWriter, data interface{}) {
	renderJSON(w, data, http.StatusCreated)
}

func HttpResponseInternalServerError(w http.ResponseWriter, data interface{}) {
	renderJSON(w, data, http.StatusInternalServerError)
}

func HttpResponseBadRequest(w http.ResponseWriter, data interface{}) {
	renderJSON(w, data, http.StatusBadRequest)
}
