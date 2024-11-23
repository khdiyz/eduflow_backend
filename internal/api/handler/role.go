package handler

import (
	"eduflow/internal/api/response"
	"eduflow/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Description Create Role
// @Summary Create Role
// @Tags Role
// @Accept json
// @Produce json
// @Param create body models.CreateRole true "Create Role"
// @Success 201 {object} response.BaseResponse
// @Failure 400,401,404,500 {object} response.BaseResponse
// @Router /api/v1/roles [post]
// @Security ApiKeyAuth
func (h *Handler) createRole(c *gin.Context) {
	var body models.CreateRole

	if err := c.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	err := h.service.Role.Create(body)
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.BaseResponse{
		Message: response.CreatedMessage,
	})
}
