package handler

import (
	"eduflow/internal/model"
	"eduflow/pkg/response"
	"eduflow/pkg/validator"
	"errors"

	"github.com/gin-gonic/gin"
)

// Login
// @Description Login User
// @Summary Login User
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body model.LoginRequest true "Login"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /auth/login [post]
func (h *Handler) login(c *gin.Context) {
	var input model.LoginRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err := validator.ValidatePayloads(input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	accessToken, refreshToken, err := h.services.Authorization.Login(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, gin.H{
		"accessToken":  accessToken.Token,
		"refreshToken": refreshToken.Token,
	}, nil)
}

// Refresh token
// @Description Refresh Token
// @Summary Refresh Token
// @Tags Auth
// @Accept json
// @Produce json
// @Param token body model.RefreshRequest true "Refresh Token"
// @Success 200 {object} model.BaseResponse
// @Failure 400 {object} model.BaseResponse
// @Failure 404 {object} model.BaseResponse
// @Failure 500 {object} model.BaseResponse
// @Router /auth/refresh [post]
func (h *Handler) refresh(c *gin.Context) {
	var (
		err   error
		input model.RefreshRequest
	)

	if err = c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	if err = validator.ValidatePayloads(input); err != nil {
		response.ErrorResponse(c, response.BadRequest, err)
		return
	}

	claims, err := h.services.Authorization.ParseToken(input.Token)
	if err != nil {
		response.AbortResponse(c, err.Error())
		return
	}

	if claims.Type != "refresh" {
		response.ErrorResponse(c, response.BadRequest, errors.New("token type must be refresh"))
		return
	}

	user, err := h.services.User.GetById(claims.UserId)
	if err != nil {
		response.FromError(c, err)
		return
	}

	accessToken, refreshToken, err := h.services.Authorization.GenerateTokens(user)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.SuccessResponse(c, response.OK, gin.H{
		"accessToken":  accessToken.Token,
		"refreshToken": refreshToken.Token,
	}, nil)
}
