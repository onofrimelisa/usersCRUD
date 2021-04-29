package service

import (
	"errors"
	"github.com/onofrimelisa/usersCRUD/internal/models"
	"github.com/onofrimelisa/usersCRUD/internal/repository"
)

type UserService interface {
	Create(u *models.BaseUser) error
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