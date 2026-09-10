package http

import "github.com/labstack/echo/v4"

func RegisterRoutes(
	e *echo.Echo,
	handler *Handler,
) {

	v1 := e.Group("/api/v1")

	transactions := v1.Group("/transactions")

	transactions.POST("/topup", handler.CreateTopUp)
	transactions.POST("/pay", handler.CreatePayment)
	transactions.GET("", handler.GetAll)
	transactions.GET("/:id", handler.GetByID)
}
