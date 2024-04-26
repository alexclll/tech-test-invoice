package business

import "test-tech-invoice/src/invoice/business/entity"

type InvoiceRepository interface {
	GetById(id int) (*entity.Invoice, error)
	Save(invoice *entity.Invoice) error
}
