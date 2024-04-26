package business

import "errors"

var ErrUserNotFound = errors.New("user_not_found")
var ErrInvoiceAlreadyPaid = errors.New("invoice_already_paid")
var ErrInvalidInvoiceAmount = errors.New("invalid_invoice_amount")
var ErrInvoiceNotFound = errors.New("invoice_not_found")
