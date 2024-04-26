package validateInvoice

import (
	"test-tech-invoice/src/invoice/business"
)

type Service struct {
	invoiceRepository business.InvoiceRepository
	userRepository    business.UserRepository
}

func (service Service) Execute(query Query) error {
	invoice, storageError := service.invoiceRepository.GetById(query.InvoiceId)

	if storageError != nil {
		return business.ErrInvoiceNotFound
	}

	if invoice.Amount != query.Amount {
		return business.ErrInvalidInvoiceAmount
	}

	if invoice.IsAlreadyPaid() {
		return business.ErrInvoiceAlreadyPaid
	}

	invoice.SetPaidStatus()

	storageError = service.invoiceRepository.Save(invoice)

	if storageError != nil {
		return storageError
	}

	user, storageError := service.userRepository.GetById(invoice.UserId)

	if storageError != nil {
		return business.ErrUserNotFound
	}

	user.AddToBalance(query.Amount)
	service.userRepository.Save(user)

	return nil
}

func NewService(invoiceRepository business.InvoiceRepository, userRepository business.UserRepository) Service {
	return Service{
		invoiceRepository: invoiceRepository,
		userRepository:    userRepository,
	}
}
