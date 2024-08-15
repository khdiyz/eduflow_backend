package handler

import (
	"eduflow/docs"
	"eduflow/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.HandleMethodNotAllowed = true
	router.Use(corsMiddleware())

	//swagger settings
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler), func(c *gin.Context) {
		docs.SwaggerInfo.Host = c.Request.Host
		if c.Request.TLS != nil {
			docs.SwaggerInfo.Schemes = []string{"https"}
		}
	})

	// AUTH
	router.POST("/auth/login", h.login)
	router.POST("/auth/refresh", h.refresh)

	api := router.Group("/api", h.userIdentity)
	{
		v1 := api.Group("/v1")
		{
			// Minio API
			v1.POST("/upload-image", h.uploadImage)

			// Role API
			role := v1.Group("/roles")
			role.GET("", h.getListRole)
			role.GET("/:id", h.getRoleById)

			// User API
			user := v1.Group("/users")
			user.POST("", h.createUser)
			user.GET("", h.getListUser)
			user.GET("/:id", h.getUserById)
			user.PUT("/:id", h.updateUser)
			user.DELETE("/:id", h.deleteUser)

			// Course API
			course := v1.Group("/courses")
			course.POST("", h.createCourse)
			course.GET("", h.getListCourse)
			course.GET("/:id", h.getCourseById)
			course.PUT("/:id", h.updateCourse)
			course.DELETE("/:id", h.deleteCourse)

			// Group API
			group := v1.Group("/groups")
			group.POST("", h.createGroup)
			group.GET("", h.getListGroup)
			group.GET("/:id", h.getGroupById)
			group.PUT("/:id", h.updateGroup)
			group.DELETE("/:id", h.deleteGroup)

			// Student API
			student := v1.Group("/students")
			student.POST("", h.createStudent)
			student.GET("", h.getListStudent)
		}
	}

	return router
}
