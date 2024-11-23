package models

type LoginRequest struct {
	Username string `json:"username" binding:"required" default:"superadmin"`
	Password string `json:"password" binding:"required" default:"$uper@Adm1n"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
