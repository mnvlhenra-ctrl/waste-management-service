package http

import "github.com/labstack/echo/v4"

func RegisterPickupRoutes(
	e *echo.Echo,
	handler *PickupHandler,
) {
	v1 := e.Group("/api/v1")

	pickups := v1.Group("/pickups")

	// Create pickup schedule
	pickups.POST("", handler.CreatePickup)

	// Get pickup schedule for tomorrow
	pickups.GET("/tomorrow", handler.GetTomorrowPickup)

	// Manual reminder test
	pickups.POST("/reminder-test", handler.SendReminderTest)
}
