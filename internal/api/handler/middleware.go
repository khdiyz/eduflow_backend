package handler

import (
	"eduflow/config"
	"eduflow/internal/api/response"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader = "Authorization"
	UserCtx             = "user_id"
	RoleCtx             = "role_id"
)

var (
	errUnauthorized = errors.New("unauthorized")
)

func (h *Handler) userIdentity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader(AuthorizationHeader)
		if header == "" {
			response.ErrorResponse(ctx, http.StatusUnauthorized, errUnauthorized)
			ctx.Abort()
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			response.ErrorResponse(ctx, http.StatusUnauthorized, errUnauthorized)
			ctx.Abort()
			return
		}

		if len(headerParts[1]) == 0 {
			response.ErrorResponse(ctx, http.StatusUnauthorized, errUnauthorized)
			ctx.Abort()
			return
		}

		claims, err := h.service.Authorization.ParseToken(headerParts[1])
		if err != nil {
			response.ErrorResponse(ctx, http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}

		if claims.Type != config.TokenTypeAccess {
			response.ErrorResponse(ctx, http.StatusUnauthorized, errUnauthorized)
			ctx.Abort()
			return
		}

		ctx.Set(UserCtx, claims.UserId)
		ctx.Set(RoleCtx, claims.RoleId)
		ctx.Next()
	}
}
