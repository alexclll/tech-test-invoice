package createInvoice

import (
	"test-tech-invoice/src/invoice/business"
	"test-tech-invoice/src/invoice/business/entity"
)

type Service struct {
	invoiceRepository business.InvoiceRepository
	userRepository    business.UserRepository
}

func (service Service) Execute(query Query) error {
	userExists, storageError := service.userRepository.Exists(query.UserId)

	if storageError != nil {
		return storageError
	}

	if !userExists {
		return business.ErrUserNotFound
	}

	invoice := entity.Invoice{
		Amount: query.Amount,
		Label:  query.Label,
		UserId: query.UserId,
		Status: entity.InvoicePending,
	}

	storageError = service.invoiceRepository.Save(&invoice)

	if storageError != nil {
		return storageError
	}

	return nil
}

func NewService(invoiceRepository business.InvoiceRepository, userRepository business.UserRepository) Service {
	return Service{
		invoiceRepository: invoiceRepository,
		userRepository:    userRepository,
	}
}
