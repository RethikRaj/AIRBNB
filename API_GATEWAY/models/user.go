package models

import "time"

type User struct {
	ID            int
	Name          string
	Email         string
	Password_hash string
	Created_at    time.Time
	Updated_at    time.Time
}
