package mainStorage

import (
	"test-tech-invoice/src/invoice/business/entity"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func (repository GormUserRepository) Exists(id int) (bool, error) {
	var count int64
	result := repository.db.Table("users").Where("id = ?", id).Count(&count)

	if result.Error != nil {
		return false, result.Error
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (repository GormUserRepository) GetById(id int) (*entity.User, error) {
	var user entity.User

	result := repository.db.Find(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repository GormUserRepository) Save(user *entity.User) error {
	result := repository.db.Save(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func NewUserRepository(db *gorm.DB) GormUserRepository {
	return GormUserRepository{
		db: db,
	}
}
