package handlers

import (
	"net/http"

	"github.com/RethikRaj/AIRBNB/API_GATEWAY/services"
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
	uh.userService.CreateUser()
	w.Write([]byte("Registered User"))
}
