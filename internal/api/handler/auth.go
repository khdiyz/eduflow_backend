package handler

import (
	"eduflow/internal/api/response"
	"eduflow/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Description Login User
// @Summary Login User
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login User"
// @Success 200 {object} models.LoginResponse
// @Failure 400,401,404,500 {object} response.BaseResponse
// @Router /api/v1/auth/login [post]
func (h *Handler) loginUser(c *gin.Context) {
	var body models.LoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	accessToken, refreshToken, err := h.service.Authorization.Login(body)
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.LoginResponse{
		AccessToken:  accessToken.Token,
		RefreshToken: refreshToken.Token,
	})
}
