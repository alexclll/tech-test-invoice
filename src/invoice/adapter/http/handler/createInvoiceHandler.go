package handler

import (
	"encoding/json"
	"net/http"
	"test-tech-invoice/src/framework/adapter/http/response"
	"test-tech-invoice/src/invoice/business"
	"test-tech-invoice/src/invoice/useCase/createInvoice"
)

type CreateInvoiceHandler struct {
	service createInvoice.Service
}

type CreateInvoicePayload struct {
	User_id int
	Amount  float32
	Label   string
}

func (handler *CreateInvoiceHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.SendMethodNotAllowedResponse(rw)
		return
	}

	rw.Header().Set("Content-Type", "application/json")

	var payload CreateInvoicePayload
	jsonError := json.NewDecoder(r.Body).Decode(&payload)
	if jsonError != nil {
		response.SendBadRequestError(rw, jsonError)
		return
	}

	serviceError := handler.service.Execute(
		createInvoice.Query{
			Amount: payload.Amount,
			Label:  payload.Label,
			UserId: payload.User_id,
		},
	)

	if serviceError == business.ErrUserNotFound {
		response.SendNotFoundError(rw, serviceError)
		return
	}

	if serviceError != nil {
		response.SendInternalServerError(rw, serviceError)
		return
	}

	response.SendNoContentSuccess(rw)
}

func (handler CreateInvoiceHandler) Pattern() string {
	return "/invoice"
}

func NewCreateInvoiceHandler(service createInvoice.Service) *CreateInvoiceHandler {
	return &CreateInvoiceHandler{
		service: service,
	}
}
