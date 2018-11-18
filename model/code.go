package model

import (
	"time"
)

type VoteCode struct {
	ID            int
	Code          string
	TransactionID string
	ProductID     int
	Creator       string
	CustomerID    *int
	Points        *int
	ClaimDate     *time.Time
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}

func (VoteCode) TableName() string {
	return "votecodes"
}
