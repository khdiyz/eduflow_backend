package handler

import (
	"eduflow/internal/model"
	"eduflow/pkg/helper"
	"eduflow/pkg/response"
	"eduflow/pkg/validator"

	"github.com/gin-gonic/gin"
)

// Create Student
// @Description Create Student
// @Summary Create Student
// @Tags Student
// @Accept json
// @Produce json
// @Param create body model.StudentCreateRequest true "Create Student"
// @Success 200 {object} model.BaseResponse
// @Router /api/v1/students [post]
// @Security ApiKeyAuth
func (h *Handler) createStudent(c *gin.Context) {
	var (
		err   error
		input model.StudentCreateRequest
	)

	if err = c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	id, err := h.services.Student.Create(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.Created, id, nil)
}

// Get List Student
// @Description Get List Student
// @Summary Get List Student
// @Tags Student
// @Accept json
// @Produce json
// @Param page_size query int64 true "page size" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} model.BaseResponse
// @Router /api/v1/students [get]
// @Security ApiKeyAuth
func (h *Handler) getListStudent(c *gin.Context) {
	pagination, err := helper.ListPagination(c)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	students, err := h.services.Student.GetList(&pagination)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, students, &pagination)
}
