package models

import "github.com/google/uuid"

type Role struct {
	Id          uuid.UUID        `json:"id"`
	Name        NameTranslations `json:"name"`
	Description NameTranslations `json:"description"`
}

type CreateRole struct {
	Name        NameTranslations `json:"name"`
	Description NameTranslations `json:"description"`
}

type UpdateRole struct {
	Id          uuid.UUID        `json:"-"`
	Name        NameTranslations `json:"name"`
	Description NameTranslations `json:"description"`
}
