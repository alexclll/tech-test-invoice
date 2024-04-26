package mainStorage

import (
	"test-tech-invoice/src/user/useCase/getUsers/model"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func (repository GormUserRepository) GetUsers() ([]model.User, error) {
	var users []model.User
	result := repository.db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func NewGetUsersRepository(db *gorm.DB) GormUserRepository {
	return GormUserRepository{
		db: db,
	}
}
