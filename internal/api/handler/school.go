package handler

import (
	"eduflow/internal/api/response"
	"eduflow/internal/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Description Create School
// @Summary Create School
// @Tags School
// @Accept json
// @Produce json
// @Param create body models.CreateSchool true "Create School"
// @Success 201 {object} response.IdResponse
// @Failure 400,401,404,500 {object} response.BaseResponse
// @Router /api/v1/schools [post]
// @Security ApiKeyAuth
func (h *Handler) createSchool(c *gin.Context) {
	var body models.CreateSchool

	if err := c.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	id, err := h.service.School.Create(body)
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.IdResponse{
		Id: id.String(),
	})
}

type listSchool struct {
	Data []models.School   `json:"data"`
	Meta models.Pagination `json:"meta"`
}

// @Description Get List School
// @Summary Get List School
// @Tags School
// @Accept json
// @Produce json
// @Param page query int true "page" default(1)
// @Param limit query int true "limit" default(10)
// @Param search query string false "search school"
// @Param status query bool false "status"
// @Success 200 {object} listSchool
// @Failure 400,401,404,500 {object} response.BaseResponse
// @Router /api/v1/schools [get]
// @Security ApiKeyAuth
func (h *Handler) getListSchool(c *gin.Context) {
	pagination, err := listPagination(c)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	var filter models.SchoolFilter
	filter.Limit = pagination.Limit
	filter.Offset = pagination.Offset

	search := c.Query("search")
	if search != "" {
		filter.Search = search
	}

	status := c.Query("status")
	if status != "" {
		filterStatus, err := strconv.ParseBool(status)
		if err != nil {
			response.ErrorResponse(c, http.StatusBadRequest, err)
			return
		}

		filter.Status = &filterStatus
	}

	schools, total, err := h.service.School.GetListSchool(filter)
	if err != nil {
		response.FromError(c, err)
		return
	}
	pagination.TotalCount = total

	pageCount := math.Ceil(float64(total) / float64(pagination.Limit))
	pagination.PageCount = int(pageCount)

	c.JSON(http.StatusOK, listSchool{
		Data: schools,
		Meta: pagination,
	})
}

// @Description Get School
// @Summary Get School
// @Tags School
// @Accept json
// @Produce json
// @Param id path string true "school id"
// @Success 200 {object} models.School
// @Failure 400,401,404,500 {object} response.BaseResponse
// @Router /api/v1/schools/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) getSchoolById(c *gin.Context) {

	c.JSON(http.StatusOK, models.School{})
}

// @Description Update School
// @Summary Update School
// @Tags School
// @Accept json
// @Produce json
// @Param id path string true "school id"
// @Param update body models.UpdateSchool true "Update School body"
// @Success 200 {object} response.BaseResponse
// @Failure 400,401,404,500 {object} response.BaseResponse
// @Router /api/v1/schools/{id} [put]
// @Security ApiKeyAuth
func (h *Handler) updateSchool(c *gin.Context) {

	c.JSON(http.StatusOK, response.BaseResponse{
		Message: "updated",
	})
}

// @Description Delete School
// @Summary Delete School
// @Tags School
// @Accept json
// @Produce json
// @Param id path string true "school id"
// @Success 200 {object} response.BaseResponse
// @Failure 400,401,404,500 {object} response.BaseResponse
// @Router /api/v1/schools/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) deleteSchool(c *gin.Context) {

	c.JSON(http.StatusOK, response.BaseResponse{
		Message: "deleted",
	})
}
