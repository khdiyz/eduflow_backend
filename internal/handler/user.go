package handler

import (
	"eduflow/internal/constants"
	"eduflow/internal/model"
	"eduflow/pkg/helper"
	"eduflow/pkg/response"
	"eduflow/pkg/validator"
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	roleIdQuery = "role_id"
	idQuery     = "id"
)

// Create User
// @Description Create User
// @Summary Create User
// @Tags User
// @Accept json
// @Produce json
// @Param create body model.UserCreateRequest true "Create User"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/v1/users [post]
// @Security ApiKeyAuth
func (h *Handler) createUser(c *gin.Context) {
	var (
		err   error
		input model.UserCreateRequest
	)
	if err = c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	// validate request body
	if err := validator.ValidatePayloads(input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	// check role name
	role, err := h.services.Role.GetById(input.RoleId)
	if err != nil {
		response.FromError(c, err)
		return
	}
	if role.Name == constants.RoleSuperAdmin {
		response.ErrorResponse(c, response.BadRequest, errors.New("super admin is not possible"))
		return
	}

	// create user
	id, err := h.services.User.Create(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.Created, id, nil)
}

// Get List User
// @Description Get List User
// @Summary Get List User
// @Tags User
// @Accept json
// @Produce json
// @Param role_id query int64 false "filter by role id"
// @Param page_size query int64 true "page size" default(10)
// @Param page  query int64 true "page" default(1)
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/v1/users [get]
// @Security ApiKeyAuth
func (h *Handler) getListUser(c *gin.Context) {
	pagination, err := helper.ListPagination(c)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	filter := make(map[string]interface{})

	roleId, err := helper.GetInt64Query(c, roleIdQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}
	if roleId != 0 {
		filter[roleIdQuery] = roleId
	}

	users, err := h.services.User.GetList(&pagination, filter)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, users, &pagination)
}

// Get User By Id
// @Description Get User By Id
// @Summary Get User By Id
// @Tags User
// @Accept json
// @Produce json
// @Param id path int64 true "User Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/v1/users/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) getUserById(c *gin.Context) {
	id, err := helper.GetInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	user, err := h.services.User.GetById(id)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, user, nil)
}

// Update User
// @Description Update User
// @Summary Update User
// @Tags User
// @Accept json
// @Produce json
// @Param id path int64 true "User Id"
// @Param update body model.UserUpdateRequest true "Update User"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/v1/users/{id} [put]
// @Security ApiKeyAuth
func (h *Handler) updateUser(c *gin.Context) {
	var input model.UserUpdateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	// check permission for update user
	userRole, err := helper.GetUserRole(c)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}
	if userRole != constants.RoleSuperAdmin {
		response.ErrorResponse(c, response.PermissionDenied, constants.ErrPermissionDenied)
		return
	}

	// validate request body
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

	// check role name
	role, err := h.services.Role.GetById(input.RoleId)
	if err != nil {
		response.FromError(c, err)
		return
	}
	if role.Name == constants.RoleSuperAdmin {
		response.ErrorResponse(c, response.BadRequest, errors.New("super admin is not possible"))
		return
	}

	// update user
	err = h.services.User.Update(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, nil, nil)
}

// Delete User
// @Description Delete User
// @Summary Delete User
// @Tags User
// @Accept json
// @Produce json
// @Param id path int64 true "User Id"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /api/v1/users/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) deleteUser(c *gin.Context) {
	id, err := helper.GetInt64Param(c, idQuery)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	// check permission for delete user
	userRole, err := helper.GetUserRole(c)
	if err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}
	if userRole != constants.RoleSuperAdmin {
		response.ErrorResponse(c, response.PermissionDenied, constants.ErrPermissionDenied)
		return
	}

	// check role for delete
	user, err := h.services.User.GetById(id)
	if err != nil {
		response.FromError(c, err)
		return
	}
	if user.RoleName == constants.RoleSuperAdmin {
		response.ErrorResponse(c, response.BadRequest, errors.New("super admin role cannot be deleted"))
		return
	}

	// delete user
	err = h.services.User.DeleteById(id)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, nil, nil)
}
