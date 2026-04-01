package db

import (
	"context"
	"fmt"

	"github.com/RethikRaj/AIRBNB/API_GATEWAY/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupDB(dbConfig config.DBConfig) (*pgxpool.Pool, error) {
	connStrURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	// Create a connection pool
	pool, err := pgxpool.New(context.Background(), connStrURL)
	if err != nil {
		fmt.Println("Error creating pool : ", err)
		return nil, err
	}

	// Verify the connection
	if err := pool.Ping(context.Background()); err != nil {
		fmt.Println("Error pinging database: ", err)
		return nil, err
	}

	fmt.Println("Successfully connected and pinged database")
	return pool, nil
}

// func CreateTables(pool *pgxpool.Pool) error {
// 	createTableQuery := `
// 	CREATE TABLE IF NOT EXISTS users (
// 		id SERIAL PRIMARY KEY,
// 		name TEXT NOT NULL,
// 		email TEXT UNIQUE NOT NULL,
// 		password_hash TEXT NOT NULL,
// 		created_at TIMESTAMPTZ DEFAULT NOW()
// 	);
// 	`

// 	_, err := pool.Exec(context.Background(), createTableQuery)

// 	return err
// }
