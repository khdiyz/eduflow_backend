package handler

import (
	"eduflow/pkg/helper"
	"eduflow/pkg/response"

	"github.com/gin-gonic/gin"
)

// GetListRole
// @Description Get List Role
// @Summary Get List Role
// @Tags Role
// @Accept json
// @Produce json
// @Param page_size query int64 true "page size" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/v1/roles [get]
// @Security ApiKeyAuth
func (h *Handler) getListRole(c *gin.Context) {
	pagination, err := helper.ListPagination(c)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	roles, err := h.services.Role.GetList(&pagination)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, roles, &pagination)
}

// Get Role By Id
// @Description Get Role By Id
// @Summary Get Role By Id
// @Tags Role
// @Accept json
// @Produce json
// @Param id path int64 true "Role Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/v1/roles/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) getRoleById(c *gin.Context) {
	id, err := helper.GetInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	role, err := h.services.Role.GetById(id)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, role, nil)
}