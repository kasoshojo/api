package model

import (
	"time"
)

type Product struct {
	ID        int
	Points    int
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (Product) TableName() string {
	return "products"
}
