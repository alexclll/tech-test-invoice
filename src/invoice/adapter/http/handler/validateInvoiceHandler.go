package handler

import (
	"encoding/json"
	"net/http"
	"test-tech-invoice/src/framework/adapter/http/response"
	"test-tech-invoice/src/invoice/business"
	"test-tech-invoice/src/invoice/useCase/validateInvoice"
)

type ValidateInvoiceHandler struct {
	service validateInvoice.Service
}

type ValidateInvoicePayload struct {
	Invoice_id int
	Amount     float32
}

func (handler *ValidateInvoiceHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.SendMethodNotAllowedResponse(rw)
		return
	}

	rw.Header().Set("Content-Type", "application/json")

	var payload ValidateInvoicePayload
	jsonError := json.NewDecoder(r.Body).Decode(&payload)
	if jsonError != nil {
		response.SendBadRequestError(rw, jsonError)
		return
	}

	serviceError := handler.service.Execute(
		validateInvoice.Query{
			InvoiceId: payload.Invoice_id,
			Amount:    payload.Amount,
		},
	)

	if serviceError == business.ErrInvoiceNotFound {
		response.SendNotFoundError(rw, serviceError)
		return
	}

	if serviceError == business.ErrInvalidInvoiceAmount {
		response.SendBadRequestError(rw, serviceError)
		return
	}

	if serviceError == business.ErrInvoiceAlreadyPaid {
		response.SendUnprocessableEntityError(rw, serviceError)
		return
	}

	if serviceError != nil {
		response.SendInternalServerError(rw, serviceError)
		return
	}

	response.SendNoContentSuccess(rw)
}

func (handler ValidateInvoiceHandler) Pattern() string {
	return "/transaction"
}

func NewValidateInvoiceHandler(service validateInvoice.Service) *ValidateInvoiceHandler {
	return &ValidateInvoiceHandler{
		service: service,
	}
}
