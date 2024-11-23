package handler

import (
	"eduflow/internal/api/response"
	"eduflow/internal/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Description Create Branch
// @Summary Create Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param id path string true "School Id"
// @Param create body models.CreateBranch true "Create Branch"
// @Success 201 {object} response.IdResponse
// @Failure 400,401,404,500 {object} response.BaseResponse
// @Router /api/v1/schools/{id}/branches [post]
// @Security ApiKeyAuth
func (h *Handler) createBranch(c *gin.Context) {
	schoolId, err := getUUIDParam(c, "id")
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	var body models.CreateBranch

	if err := c.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	body.SchoolId = schoolId

	id, err := h.service.Branch.CreateBranch(body)
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.IdResponse{
		Id: id.String(),
	})
}

type getBranchesResponse struct {
	Data []models.Branch   `json:"data"`
	Meta models.Pagination `json:"meta"`
}

// @Description Get List Branch
// @Summary Get List Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param id path string true "School Id"
// @Param page query int true "page" default(1)
// @Param limit query int true "limit" default(10)
// @Param search query string false "search"
// @Param status query bool false "status"
// @Success 200 {object} getBranchesResponse
// @Failure 400,401,404,500 {object} response.BaseResponse
// @Router /api/v1/schools/{id}/branches [get]
// @Security ApiKeyAuth
func (h *Handler) getBranches(c *gin.Context) {
	schoolId, err := getUUIDParam(c, "id")
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	pagination, err := listPagination(c)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	var filter models.BranchFilter
	filter.Limit = pagination.Limit
	filter.Offset = pagination.Offset
	filter.SchoolId = schoolId

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

	branches, total, err := h.service.Branch.GetBranches(filter)
	if err != nil {
		response.FromError(c, err)
		return
	}
	pagination.TotalCount = total

	pageCount := math.Ceil(float64(total) / float64(pagination.Limit))
	pagination.PageCount = int(pageCount)

	c.JSON(http.StatusOK, getBranchesResponse{
		Data: branches,
		Meta: pagination,
	})
}

// @Description Get Branch
// @Summary Get Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param id path string true "School id"
// @Param branch-id path string true "Branch id"
// @Success 200 {object} models.Branch
// @Failure 400,401,404,500 {object} response.BaseResponse
// @Router /api/v1/schools/{id}/branches/{branch-id} [get]
// @Security ApiKeyAuth
func (h *Handler) getBranch(c *gin.Context) {
	schoolId, err := getUUIDParam(c, "id")
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	branchId, err := getUUIDParam(c, "branch-id")
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	branch, err := h.service.Branch.GetBranch(schoolId, branchId)
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusOK, branch)
}

// @Description Update Branch
// @Summary Update Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param id path string true "School Id"
// @Param branch-id path string true "Branch Id"
// @Param update body models.UpdateBranch true "Update Branch body"
// @Success 200 {object} response.BaseResponse
// @Failure 400,401,404,500 {object} response.BaseResponse
// @Router /api/v1/schools/{id}/branches/{branch-id} [put]
// @Security ApiKeyAuth
func (h *Handler) updateBranch(c *gin.Context) {
	schoolId, err := getUUIDParam(c, "id")
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	branchId, err := getUUIDParam(c, "branch-id")
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	var body models.UpdateBranch
	if err := c.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	body.Id = branchId
	body.SchoolId = schoolId

	err = h.service.Branch.UpdateBranch(body)
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{
		Message: "updated",
	})
}

// @Description Delete Branch
// @Summary Delete Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param id path string true "School Id"
// @Param branch-id path string true "Branch Id"
// @Success 200 {object} response.BaseResponse
// @Failure 400,401,404,500 {object} response.BaseResponse
// @Router /api/v1/schools/{id}/branches/{branch-id} [delete]
// @Security ApiKeyAuth
func (h *Handler) deleteBranch(c *gin.Context) {
	schoolId, err := getUUIDParam(c, "id")
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	branchId, err := getUUIDParam(c, "branch-id")
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	err = h.service.Branch.DeleteBranch(schoolId, branchId)
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{
		Message: "deleted",
	})
}
