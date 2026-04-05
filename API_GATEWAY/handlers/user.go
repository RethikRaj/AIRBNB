package handlers

import (
	"errors"
	"net/http"

	"github.com/RethikRaj/AIRBNB/API_GATEWAY/contextkeys"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/dto"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/services"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/utils"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(_userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: _userService,
	}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Step 1, 2 : Decode the JSON input (Deserialization) , Validation of input : Done by middleware

	// Step 3 : Call service layer
	// The below line is dangerous as it can panic if the assertion fails
	// err := uh.userService.CreateUser(r.Context().Value(contextkeys.CreateUserPayload).(*dto.CreateUserRequest))

	// 3.1) Safe assertion
	payload, ok := r.Context().Value(contextkeys.CreateUserPayload).(*dto.CreateUserRequest)
	if !ok || payload == nil {
		utils.WriteErrorJsonResponse(w, http.StatusBadRequest, "Invalid payload for create user", errors.New("Invalid Payload"))
		return
	}

	// 3.2) Call service layer
	err := uh.userService.CreateUser(payload)

	if err != nil {
		utils.WriteErrorJsonResponse(w, http.StatusInternalServerError, "Create user request failed", err)
		return
	}

	// Step 4 : Encode(Serialization) and write JSON response
	utils.WriteSuccessJsonResponse(w, http.StatusCreated, "User created succesfully", nil)
}

func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(contextkeys.UserID).(int)

	if !ok {
		utils.WriteErrorJsonResponse(w, http.StatusBadRequest, "Invalid ID", errors.New("Invalid ID"))
		return
	}

	user, err := uh.userService.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			utils.WriteErrorJsonResponse(w, http.StatusNotFound, "User not found", err)
			return
		}
		utils.WriteErrorJsonResponse(w, http.StatusInternalServerError, "Error fetching user", err)
		return
	}

	utils.WriteSuccessJsonResponse(w, http.StatusOK, "User fetched successfully", user)
}

func (uh *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	// Read payload from context
	payload, ok := r.Context().Value(contextkeys.SignInUserPayload).(*dto.SignInUserRequest)

	if !ok || payload == nil {
		utils.WriteErrorJsonResponse(w, http.StatusBadRequest, "Invalid payload for sign up user", errors.New("Invalid Payload"))
		return
	}

	signedToken, err := uh.userService.LoginUser(payload)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrNotFound):
			utils.WriteErrorJsonResponse(w, http.StatusNotFound, "No user found. Please sign up", err)
		case errors.Is(err, services.ErrInvalidCredentials):
			utils.WriteErrorJsonResponse(w, http.StatusUnauthorized, "Invalid credentials", err)
		default:
			utils.WriteErrorJsonResponse(w, http.StatusInternalServerError, "Internal server error", err)
		}
		return
	}

	utils.WriteSuccessJsonResponse(w, http.StatusOK, "User signed up successfully", signedToken)
}
