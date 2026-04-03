package dto

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=30"`
}
