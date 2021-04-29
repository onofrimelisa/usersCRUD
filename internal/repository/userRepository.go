package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/onofrimelisa/usersCRUD/internal/db"
	"github.com/onofrimelisa/usersCRUD/internal/models"
)

type UserRepository interface {
	Create(u *models.BaseUser) error
	GetUsers(u *[]models.BaseUser) error
	GetUser(u *models.BaseUser, id string) error
	UpdateUser(u *models.BaseUser) error
	DeleteUser(u *models.BaseUser) error
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Create(u *models.BaseUser) error {
	err := r.db.Create(u)

	return err.Error
}

func (r *userRepository) GetUsers(u *[]models.BaseUser) error {
	err := r.db.Find(u)

	return err.Error
}

func (r *userRepository) GetUser(u *models.BaseUser, id string) error {
	err := r.db.First(u, id)

	if u.Id == 0 {
		return errors.New("user not found")
	}

	return err.Error
}

func (r *userRepository) UpdateUser(u *models.BaseUser) error {
	err := r.db.Save(u)

	return err.Error
}

func (r *userRepository) DeleteUser(u *models.BaseUser) error {
	err := r.db.Delete(u)

	return err.Error
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: db.InitDb(),
	}
}