package http

import "github.com/labstack/echo/v4"

func RegisterRoutes(
	e *echo.Echo,
	handler *Handler,
) {

	v1 := e.Group("/api/v1")

	wastes := v1.Group("/wastes")

	wastes.POST("", handler.Create)
	wastes.GET("", handler.GetAll)
	wastes.GET("/:id", handler.GetByID)
	wastes.PUT("/:id", handler.Update)
	wastes.DELETE("/:id", handler.Delete)
}
