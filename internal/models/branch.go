package models

import (
	"time"

	"github.com/google/uuid"
)

type Branch struct {
	Id           uuid.UUID  `json:"id"`
	SchoolId     uuid.UUID  `json:"school_id"`
	Name         string     `json:"name"`
	Address      string     `json:"address"`
	Email        string     `json:"email"`
	PhoneNumber  string     `json:"phone_number"`
	OpeningHours string     `json:"opening_hours"`
	Status       bool       `json:"status"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type BranchFilter struct {
	Search   string
	Status   *bool
	SchoolId uuid.UUID
	Limit    int
	Offset   int
}

type CreateBranch struct {
	SchoolId     uuid.UUID `json:"-"`
	Name         string    `json:"name" binding:"required"`
	Address      string    `json:"address"`
	Email        string    `json:"email" binding:"omitempty,email"`
	PhoneNumber  string    `json:"phone_number"`
	OpeningHours string    `json:"opening_hours"`
	Status       bool      `json:"-"`
}

type UpdateBranch struct {
	Id           uuid.UUID `json:"-"`
	SchoolId     uuid.UUID `json:"-"`
	Name         string    `json:"name" binding:"required"`
	Address      string    `json:"address"`
	Email        string    `json:"email" binding:"omitempty,email"`
	PhoneNumber  string    `json:"phone_number"`
	OpeningHours string    `json:"opening_hours"`
	Status       bool      `json:"status"`
}
