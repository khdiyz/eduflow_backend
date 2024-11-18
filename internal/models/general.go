package models

type NameTranslations struct {
	Uz string `json:"uz" binding:"required"`
	En string `json:"en" binding:"required"`
	Ru string `json:"ru" binding:"required"`
}

type Pagination struct {
	Page       int `json:"page"  default:"1"`
	Limit      int `json:"limit" default:"10"`
	Offset     int `json:"-" default:"0"`
	PageCount  int `json:"page_count"`
	TotalCount int `json:"total_count"`
}
