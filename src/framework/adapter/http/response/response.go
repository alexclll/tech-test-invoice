package response

import (
	"encoding/json"
	"net/http"
)

func SendBadRequestError(rw http.ResponseWriter, appError error) {
	sendError(rw, appError, http.StatusBadRequest)
}

func SendInternalServerError(rw http.ResponseWriter, appError error) {
	sendError(rw, appError, http.StatusInternalServerError)
}

func SendMethodNotAllowedResponse(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func SendNoContentSuccess(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusNoContent)
}

func SendNotFoundError(rw http.ResponseWriter, appError error) {
	sendError(rw, appError, http.StatusNotFound)
}

func SendUnprocessableEntityError(rw http.ResponseWriter, appError error) {
	sendError(rw, appError, http.StatusUnprocessableEntity)
}

func sendError(rw http.ResponseWriter, appError error, statusCode int) {
	json, jsonError := json.Marshal(map[string]string{"error": appError.Error()})

	if jsonError != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(statusCode)
	rw.Write(json)
}
