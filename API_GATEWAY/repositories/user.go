package repositories

import (
	"context"
	"fmt"

	"github.com/RethikRaj/AIRBNB/API_GATEWAY/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(name string, email string, password_hash string) error
	GetUserByID(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	pool *pgxpool.Pool
}

// To achieve dependency injection , in the below function why do we need to return an interface, when we can also achieve dependency injection by returning *userRepository  ?
// Way 1 : Return an *userRepository
// Problem : Callers should be disciplined enough to write the code like this :
// var repo UserRepository = NewUserRepository() // explicitly typed
// If by chance they write : userRepository := NewUserRepository() // Tight Coupling. The caller starts to depend on concrete implementation instead of interface.
//
// Therefore , we return an interface
func NewUserRepository(_pool *pgxpool.Pool) UserRepository {
	return &userRepository{
		pool: _pool,
	}
}

func (ur *userRepository) CreateUser(name string, email string, password_hash string) error {
	// Step 1 : Prepare the query
	insertUserQuery :=
		`INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3)`

	// Step 2 : Execute the query
	result, err := ur.pool.Exec(context.Background(), insertUserQuery, name, email, password_hash)

	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	fmt.Println(result.RowsAffected())

	if result.RowsAffected() == 0 {
		return fmt.Errorf("failed to insert user")
	}

	// User created succesfully
	return nil
}

func (ur *userRepository) GetUserByID(id int) (*models.User, error) {
	// Step 1 : Prepare the query
	query :=
		`SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1`

	// Step 2 : Execute the query
	row := ur.pool.QueryRow(context.Background(), query, id)

	// Step 3 : Process the result and prepare a desired output

	var user models.User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Created_at, &user.Updated_at)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) GetUserByEmail(email string) (*models.User, error) {
	query :=
		`SELECT id, name, email, password_hash FROM users WHERE email = $1`

	row := ur.pool.QueryRow(context.Background(), query, email)

	// fmt.Println(row)

	var user models.User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password_hash)

	if err != nil {
		return nil, err
	}

	return &user, nil

}
