package getUsers

import (
	"test-tech-invoice/src/user/useCase/getUsers/model"
)

type UserRepository interface {
	GetUsers() ([]model.User, error)
}
