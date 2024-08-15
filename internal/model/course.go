package model

import "time"

type Course struct {
	Id          int64      `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Photo       string     `json:"photo"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type CourseShort struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type CourseCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
}

type CourseUpdateRequest struct {
	Id          int64  `json:"-"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
}
