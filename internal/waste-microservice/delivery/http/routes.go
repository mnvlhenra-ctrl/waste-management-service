package http

import (
	"waste-management-service/internal/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(
	e *echo.Echo,
	handler *Handler,
) {
	v1 := e.Group("/api/v1")

	wastes := v1.Group("/wastes")

	wastes.Use(middleware.JWTMiddleware())

	wastes.GET("", handler.GetAll)
	wastes.GET("/:id", handler.GetByID)

	wastes.POST("", handler.Create, middleware.AdminOnly)
	wastes.PUT("/:id", handler.Update, middleware.AdminOnly)
	wastes.DELETE("/:id", handler.Delete, middleware.AdminOnly)
}
