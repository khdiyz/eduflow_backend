package model

import "time"

type User struct {
	Id          int64     `json:"id"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone"`
	BirthDate   time.Time `json:"birth_date"`
	Photo       string    `json:"photo"`
	Username    string    `json:"username"`
	Password    string    `json:"-"`

	RoleId   int64  `json:"role_id"`
	RoleName string `json:"role_name"`

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserShort struct {
	Id       int64  `json:"id"`
	FullName string `json:"full_name"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required" default:"superadmin"`
	Password string `json:"password" validate:"required" default:"P@ssw0rd2o24"`
}

type RefreshRequest struct {
	Token string `json:"token" validate:"required"`
}

type UserCreateRequest struct {
	FullName    string    `json:"full_name" validate:"required"`
	PhoneNumber string    `json:"phone" validate:"uzbphone"`
	BirthDate   time.Time `json:"birth_date"`
	Photo       string    `json:"photo"`
	RoleId      int64     `json:"role_id"`
	Username    string    `json:"username" validate:"required"`
	Password    string    `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Id          int64     `json:"-"`
	FullName    string    `json:"full_name" validate:"required"`
	PhoneNumber string    `json:"phone" validate:"uzbphone"`
	BirthDate   time.Time `json:"birth_date"`
	Photo       string    `json:"photo"`
	RoleId      int64     `json:"role_id"`
	Username    string    `json:"username" validate:"required"`
}
