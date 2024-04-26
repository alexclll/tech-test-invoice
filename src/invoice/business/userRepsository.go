package business

import (
	"test-tech-invoice/src/invoice/business/entity"
)

type UserRepository interface {
	GetById(id int) (*entity.User, error)
	Exists(id int) (bool, error)
	Save(user *entity.User) error
}
