package model

import (
	"time"

	"github.com/kasoshojo/api/app"
)

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

func (user *User) ToAppUser() app.GoaUser {
	ret := app.GoaUser{}
	ret.ID = user.ID
	ret.Birthdate = &user.Birthday
	ret.CreatedAt = &user.CreatedAt
	ret.GivenNames = &user.GivenNames
	ret.LastNames = &user.LastNames
	ret.UpdatedAt = user.UpdatedAt
	ret.Location = &user.Location
	ret.Username = user.Username
	ret.Points = &user.Points
	return ret
}
