package model

import "time"

type Student struct {
	Id        int64      `json:"id"`
	FullName  string     `json:"full_name"`
	Phone1    string     `json:"phone_1"`
	Phone2    string     `json:"phone_2"`
	Address   string     `json:"address"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type StudentCreateRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Phone1   string `json:"phone_1" validate:"uzbphone,required"`
	Phone2   string `json:"phone_2" validate:"uzbphone"`
	Address  string `json:"address"`
}
