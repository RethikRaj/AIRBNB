package services

import (
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/models"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/repositories"
)

type UserService interface {
	CreateUser() error
	GetUserByID(id int) (*models.User, error)
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
	us.userRepository.CreateUser("test_user", "test_user@gmail.com", "123456789")
	return nil
}

func (us *userService) GetUserByID(id int) (*models.User, error) {
	user, err := us.userRepository.GetUserByID(id)
	return user, err
}
