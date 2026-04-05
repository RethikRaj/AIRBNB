package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RethikRaj/AIRBNB/API_GATEWAY/contextkeys"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/dto"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/services"
	"github.com/RethikRaj/AIRBNB/API_GATEWAY/utils"
	"github.com/go-chi/chi/v5"
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
	}

	// 3.2) Call service layer
	err := uh.userService.CreateUser(payload)

	if err != nil {
		utils.WriteErrorJsonResponse(w, http.StatusInternalServerError, "Create user request failed", err)
		return
	}

	// Step 4 : Encode(Serialization) and write JSON response
	utils.WriteSuccessJsonResponse(w, http.StatusOK, "User created succesfully", nil)
}

func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user, err := uh.userService.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	fmt.Println("Fetched user :", user)
	w.Write([]byte("User fetched succesfully"))
}

func (uh *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	signedToken, err := uh.userService.LoginUser()
	if err != nil {
		http.Error(w, "Invalid Credentials", http.StatusBadRequest)
	}
	w.Write([]byte(signedToken))
}
