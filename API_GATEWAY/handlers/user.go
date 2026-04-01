package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/RethikRaj/AIRBNB/API_GATEWAY/services"
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
	uh.userService.CreateUser()
	w.Write([]byte("Registered User"))
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
