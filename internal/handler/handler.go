package handler

import (
	"eduflow/config"
	"eduflow/docs"
	"eduflow/internal/service"
	"eduflow/pkg/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
}

func NewHandler(service *service.Service, loggers *logger.Logger) *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.HandleMethodNotAllowed = true
	// router.Use(corsMiddleware())

	//swagger settings
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler), func(ctx *gin.Context) {
		docs.SwaggerInfo.Host = ctx.Request.Host
		if ctx.Request.TLS != nil {
			docs.SwaggerInfo.Schemes = []string{"https"}
		}
	})

	return router
}
