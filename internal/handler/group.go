package handler

import (
	"eduflow/internal/model"
	"eduflow/pkg/helper"
	"eduflow/pkg/response"
	"eduflow/pkg/validator"

	"github.com/gin-gonic/gin"
)

// Create Group
// @Description Create Group
// @Summary Create Group
// @Tags Group
// @Accept json
// @Produce json
// @Param create body model.GroupCreateRequest true "Create Group"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/v1/groups [post]
// @Security ApiKeyAuth
func (h *Handler) createGroup(c *gin.Context) {
	var (
		err   error
		input model.GroupCreateRequest
	)

	if err = c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	id, err := h.services.Group.Create(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.Created, id, nil)
}

// Get List Group
// @Description Get List Group
// @Summary Get List Group
// @Tags Group
// @Accept json
// @Produce json
// @Param page_size query int64 true "page size" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/v1/groups [get]
// @Security ApiKeyAuth
func (h *Handler) getListGroup(c *gin.Context) {
	pagination, err := helper.ListPagination(c)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	groups, err := h.services.Group.GetList(&pagination)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, groups, &pagination)
}

// Get Group By Id
// @Description Get Group By Id
// @Summary Get Group By Id
// @Tags Group
// @Accept json
// @Produce json
// @Param id path int64 true "Group Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/v1/groups/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) getGroupById(c *gin.Context) {
	id, err := helper.GetInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	group, err := h.services.Group.GetById(id)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, group, nil)
}

// Update Group
// @Description Update Group
// @Summary Update Group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path int64 true "Group Id"
// @Param update body model.GroupUpdateRequest true "Update Group"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/v1/groups/{id} [put]
// @Security ApiKeyAuth
func (h *Handler) updateGroup(c *gin.Context) {
	var input model.GroupUpdateRequest

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

	err = h.services.Group.Update(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, nil, nil)
}

// Delete Group
// @Description Delete Group
// @Summary Delete Group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path int64 true "Group Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/v1/groups/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) deleteGroup(c *gin.Context) {
	id, err := helper.GetInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	err = h.services.Group.Delete(id)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, nil, nil)
}
