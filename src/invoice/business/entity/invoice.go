package entity

const (
	InvoicePending = "pending"
	InvoicePaid    = "paid"
)

type Invoice struct {
	Id     int `gorm:"primary_key"`
	Amount float32
	Label  string
	Status string
	UserId int `gorm:"column:user_id"`
}

func (invoice *Invoice) IsAlreadyPaid() bool {
	return invoice.Status == InvoicePaid
}

func (invoice *Invoice) SetPaidStatus() {
	invoice.Status = InvoicePaid
}
