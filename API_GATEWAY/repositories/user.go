package repositories

import "fmt"

type UserRepository interface {
	CreateUser() error
}

type userRepository struct {
}

// To achieve dependency injection , in the below function why do we need to return an interface, when we can also achieve dependency injection by returning *userRepository  ?
// Way 1 : Return an *userRepository
// Problem : Callers should be disciplined enough to write the code like this :
// var repo UserRepository = NewUserRepository() // explicitly typed
// If by chance they write : userRepository := NewUserRepository() // Tight Coupling. The caller starts to depend on concrete implementation instead of interface.
//
// Therefore , we return an interface
func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (ur *userRepository) CreateUser() error {
	fmt.Println("Repository : Creating User...")
	return nil
}
