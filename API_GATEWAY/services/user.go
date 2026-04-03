package services

import (
	"fmt"

	"github.com/RethikRaj/AIRBNB/API_GATEWAY/dto"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/models"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/repositories"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/utils"
)

type UserService interface {
	CreateUser(createUserRequestPayload *dto.CreateUserRequest) error
	GetUserByID(id int) (*models.User, error)
	LoginUser() (string, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(_userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: _userRepository,
	}
}

func (us *userService) CreateUser(createUserRequestPayload *dto.CreateUserRequest) error {
	// Step 0 : Validation done in middlewares
	// Step 1 : Hash the password
	password_hash, err := utils.HashPassword(createUserRequestPayload.Password)

	if err != nil {
		return fmt.Errorf("Error while hashing the password : %w", err)
	}

	// Step 2 : Call repository layer to create user
	err = us.userRepository.CreateUser(createUserRequestPayload.Name, createUserRequestPayload.Email, password_hash)

	if err != nil {
		return fmt.Errorf("Failed to create user : %w", err)
	}

	return nil
}

func (us *userService) GetUserByID(id int) (*models.User, error) {
	user, err := us.userRepository.GetUserByID(id)
	return user, err
}

func (us *userService) LoginUser() (string, error) {
	email := "test_user@gmail.com"
	password := "1234567891"

	// step 1 : Compare password and hashed Password
	user, err := us.userRepository.GetUserByEmail(email)

	if err != nil {
		return "", err
	}
	isPasswordMatch := utils.CompareAndVerifyPassword(user.Password_hash, password)

	if !isPasswordMatch {
		return "", fmt.Errorf("Invalid credentials")
	}

	// Step 2 : Generate token
	signedToken, err := utils.CreateJWTToken(user.ID)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
