package handler

import (
	"eduflow/internal/model"
	"eduflow/pkg/helper"
	"eduflow/pkg/response"
	"eduflow/pkg/validator"

	"github.com/gin-gonic/gin"
)

// Create Course
// @Description Create Course
// @Summary Create Course
// @Tags Course
// @Accept json
// @Produce json
// @Param create body model.CourseCreateRequest true "Create Course"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/v1/courses [post]
// @Security ApiKeyAuth
func (h *Handler) createCourse(c *gin.Context) {
	var (
		err   error
		input model.CourseCreateRequest
	)

	if err = c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	id, err := h.services.Course.Create(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.Created, id, nil)
}

// Get List Course
// @Description Get List Course
// @Summary Get List Course
// @Tags Course
// @Accept json
// @Produce json
// @Param page_size query int64 true "page size" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/v1/courses [get]
// @Security ApiKeyAuth
func (h *Handler) getListCourse(c *gin.Context) {
	pagination, err := helper.ListPagination(c)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	courses, err := h.services.Course.GetList(&pagination)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, courses, &pagination)
}

// Get Course By Id
// @Description Get Course By Id
// @Summary Get Course By Id
// @Tags Course
// @Accept json
// @Produce json
// @Param id path int64 true "Course Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/v1/courses/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) getCourseById(c *gin.Context) {
	id, err := helper.GetInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	course, err := h.services.Course.GetById(id)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, course, nil)
}

// Update Course
// @Description Update Course
// @Summary Update Course
// @Tags Course
// @Accept json
// @Produce json
// @Param id path int64 true "Course Id"
// @Param update body model.CourseUpdateRequest true "Update Course"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/v1/courses/{id} [put]
// @Security ApiKeyAuth
func (h *Handler) updateCourse(c *gin.Context) {
	var input model.CourseUpdateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	id, err := helper.GetInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}
	input.Id = id

	err = h.services.Course.Update(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, nil, nil)
}

// Delete Course
// @Description Delete Course
// @Summary Delete Course
// @Tags Course
// @Accept json
// @Produce json
// @Param id path int64 true "Course Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/v1/courses/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) deleteCourse(c *gin.Context) {
	id, err := helper.GetInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	err = h.services.Course.Delete(id)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, nil, nil)
}
