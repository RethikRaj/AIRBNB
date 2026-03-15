package services

import (
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/repositories"
)

type UserService interface {
	CreateUser() error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(_userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: _userRepository,
	}
}

func (us *userService) CreateUser() error {
	us.userRepository.CreateUser()
	return nil
}
