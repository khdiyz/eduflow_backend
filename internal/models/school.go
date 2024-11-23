package models

import (
	"time"

	"github.com/google/uuid"
)

type School struct {
	Id          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Address     string     `json:"address"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phone_number"`
	Currency    string     `json:"currency"`
	Timezone    string     `json:"timezone"`
	Status      bool       `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type SchoolFilter struct {
	Search string
	Status *bool
	Limit  int
	Offset int
}

type CreateSchool struct {
	Name        string `json:"name" binding:"required"`
	Address      string `json:"address"`
	Email       string `json:"email" binding:"omitempty,email"`
	PhoneNumber string `json:"phone_number"`
	Currency    string `json:"currency" binding:"required"`
	Timezone    string `json:"timezone" binding:"required"`
	Status      bool   `json:"-"`
}

type UpdateSchool struct {
	Id          uuid.UUID `json:"-"`
	Name        string    `json:"name" binding:"required"`
	Addess      string    `json:"address"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Currency    string    `json:"currency" binding:"required"`
	Timezone    string    `json:"timezone" binding:"required"`
	Status      bool      `json:"status" binding:"required"`
}
