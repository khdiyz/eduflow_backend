package handler

import (
	"eduflow/config"
	"eduflow/docs"
	"eduflow/internal/middleware"
	"eduflow/internal/service"
	"eduflow/pkg/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandlers(service *service.Service, loggers *logger.Logger) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.HandleMethodNotAllowed = true
	router.Use(middleware.CorsMiddleware())

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler), func(ctx *gin.Context) {
		docs.SwaggerInfo.Host = ctx.Request.Host
		if ctx.Request.TLS != nil {
			docs.SwaggerInfo.Schemes = []string{"https"}
		}
	})

	v1 := router.Group("/api/v1")
	{
		roles := v1.Group("/roles")
		{
			roles.POST("", h.createRole)
		}
	}

	return router
}
