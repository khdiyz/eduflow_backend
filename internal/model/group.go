package model

import "time"

type Group struct {
	Id        int64      `json:"id"`
	Name      string     `json:"name"`
	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`

	Course  CourseShort `json:"course,omitempty"`
	Teacher UserShort   `json:"user,omitempty"`
}

type GroupCreateRequest struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	CourseId  int64     `json:"course_id"`
	TeacherId int64     `json:"teacher_id"`
}

type GroupUpdateRequest struct {
	Id        int64      `json:"-"`
	Name      string     `json:"name"`
	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	CourseId  int64      `json:"course_id"`
	TeacherId int64      `json:"teacher_id"`
}
