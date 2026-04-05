package services

import (
	"errors"
	"fmt"

	"github.com/RethikRaj/AIRBNB/API_GATEWAY/dto"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/models"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/repositories"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/utils"
	"github.com/jackc/pgx/v5"
)

type UserService interface {
	CreateUser(createUserRequestPayload *dto.CreateUserRequest) error
	GetUserByID(id int) (*models.User, error)
	LoginUser(payload *dto.SignInUserRequest) (string, error)
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

func (us *userService) LoginUser(payload *dto.SignInUserRequest) (string, error) {

	// step 1 : Compare password and hashed Password
	user, err := us.userRepository.GetUserByEmail(payload.Email)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", ErrNotFound
		}
		return "", err
	}

	isPasswordMatch := utils.CompareAndVerifyPassword(user.Password_hash, payload.Password)

	if !isPasswordMatch {
		return "", ErrInvalidCredentials
	}

	// Step 2 : Generate token
	signedToken, err := utils.CreateJWTToken(user.ID)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
