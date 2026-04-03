package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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
	// Step 1 : Decode the JSON input (Deserialization)
	var createUserRequestPayload dto.CreateUserRequest

	if err := utils.ReadJsonBody(r, &createUserRequestPayload); err != nil {
		utils.WriteErrorJsonResponse(w, http.StatusBadRequest, "Error decoding json", err)
		return
	}

	// Step 2 : Validation of input
	if err := utils.Validate.Struct(&createUserRequestPayload); err != nil {
		utils.WriteErrorJsonResponse(w, http.StatusBadRequest, "Invalid JSON", err)
		return
	}

	// Step 3 : Call service layer
	// Note : Here it is passed by reference so that we don't create unnecessary copies
	err := uh.userService.CreateUser(&createUserRequestPayload)

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
