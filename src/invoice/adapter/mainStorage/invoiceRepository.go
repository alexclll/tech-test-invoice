package mainStorage

import (
	"test-tech-invoice/src/invoice/business/entity"

	"gorm.io/gorm"
)

type GormInvoiceRepository struct {
	db *gorm.DB
}

func (repository GormInvoiceRepository) GetById(id int) (*entity.Invoice, error) {
	var invoice entity.Invoice

	result := repository.db.Find(&invoice, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &invoice, nil
}

func (repository GormInvoiceRepository) Save(invoice *entity.Invoice) error {
	result := repository.db.Save(invoice)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func NewInvoiceRepository(db *gorm.DB) GormInvoiceRepository {
	return GormInvoiceRepository{
		db: db,
	}
}
