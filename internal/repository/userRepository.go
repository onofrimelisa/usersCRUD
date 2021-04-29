package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/onofrimelisa/usersCRUD/internal/db"
	"github.com/onofrimelisa/usersCRUD/internal/models"
)

type UserRepository interface {
	Create(u *models.BaseUser) error
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Create(u *models.BaseUser) error {
	err := r.db.Create(&u)

	return err.Error
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: db.InitDb(),
	}
}