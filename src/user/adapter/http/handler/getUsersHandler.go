package handler

import (
	"encoding/json"
	"net/http"
	"test-tech-invoice/src/framework/adapter/http/response"
	"test-tech-invoice/src/user/useCase/getUsers"
)

type GetUsersHandler struct {
	service getUsers.Service
}

func (handler *GetUsersHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.SendMethodNotAllowedResponse(rw)
		return
	}

	rw.Header().Set("Content-Type", "application/json")

	serviceResponse, serviceError := handler.service.Execute()

	if serviceError != nil {
		response.SendInternalServerError(rw, serviceError)
	}

	jsonError := json.NewEncoder(rw).Encode(serviceResponse.Users)
	if jsonError != nil {
		response.SendInternalServerError(rw, jsonError)
	}
}

func (handler GetUsersHandler) Pattern() string {
	return "/users"
}

func NewGetUsersHandler(service getUsers.Service) *GetUsersHandler {
	return &GetUsersHandler{
		service: service,
	}
}
