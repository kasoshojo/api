package model

import "time"

// User is user model
type User struct {
	ID             int
	Username       string
	Password       string
	SecretQuestion string
	SecretAnswer   string
	GivenNames     string
	LastNames      string
	Email          string
	Birthday       time.Time
	Location       string
	Points         int
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

func (User) TableName() string {
	return "customers"
}
