package model

import "time"

// User is user model
type User struct {
	ID        int
	Username  string
	Password  string
	Fname     *string
	Lname     *string
	Mail      string
	Birthdate time.Time
	Location  *string
	Points    int
	CreatedAt time.Time
	UpdatedAt *time.Time
}
