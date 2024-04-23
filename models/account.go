package models

import (
	"time"
)

type Account struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"userId"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CurrencyCode string    `json:"currencyCode"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func NewAccount(userID int64, name, description, currencyCode string) *Account {
	return &Account{
		UserID:       userID,
		Name:         name,
		Description:  description,
		CurrencyCode: currencyCode,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}
}
