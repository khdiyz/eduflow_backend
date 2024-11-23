package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID  `json:"id"`
	RoleId       uuid.UUID  `json:"role_id"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	PhoneNumbers []string   `json:"phone_numbers"`
	Username     string     `json:"username"`
	Password     string     `json:"-"`
	Status       bool       `json:"status"`
	BranchId     string     `json:"branch_id,omitempty"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type UserFilter struct {
	Search   string
	RoleId   uuid.UUID
	Status   *bool
	BranchId uuid.UUID
	Limit    int
	Offset   int
}

type CreateUser struct {
	RoleId       uuid.UUID `json:"role_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	PhoneNumbers []string  `json:"phone_numbers"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	BranchId     uuid.UUID `json:"branch_id"`
}

type UpdateUser struct {
	Id           uuid.UUID `json:"-"`
	RoleId       uuid.UUID `json:"role_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	PhoneNumbers []string  `json:"phone_numbers"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	BranchId     uuid.UUID `json:"branch_id"`
}
