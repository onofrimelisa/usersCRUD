package service

import (
	"errors"
	"github.com/onofrimelisa/usersCRUD/internal/models"
	"github.com/onofrimelisa/usersCRUD/internal/repository"
)

type UserService interface {
	Create(u *models.BaseUser) error
	GetUsers(u *[]models.BaseUser) error
	GetUser(u *models.BaseUser, id string) error
	UpdateUser(oldUser models.BaseUser, newUser *models.BaseUser) error
	DeleteUser(u *models.BaseUser) error
}

type userService struct {
	userRepository repository.UserRepository
}

func (r *userService) Create(u *models.BaseUser) error {
	if isEmpty(u.Firstname, u.Lastname) {
		return errors.New("fields are empty")
	}

	return r.userRepository.Create(u)
}

func (r *userService) GetUsers(u *[]models.BaseUser) error {
	return r.userRepository.GetUsers(u)
}

func (r *userService) GetUser(u *models.BaseUser, id string) error {
	return r.userRepository.GetUser(u, id)
}

func (r *userService) UpdateUser(oldUser models.BaseUser, newUser *models.BaseUser) error {
	if isEmpty(newUser.Firstname, newUser.Lastname) {
		return errors.New("fields are empty")
	}

	newUser.Id = oldUser.Id

	return r.userRepository.UpdateUser(newUser)
}

func (r *userService) DeleteUser(u *models.BaseUser) error {
	return r.userRepository.DeleteUser(u)
}

func isEmpty(fields ...string) bool {
	for _, field := range fields {
		if field == "" {
			return true
		}
	}

	return false
}


func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}