package http

import (
	"net/http"
	"time"

	"waste-management-service/internal/waste-microservice/usecase"

	"github.com/labstack/echo/v4"
)

type PickupHandler struct {
	usecase     *usecase.PickupScheduleUsecase
	reminderJob *usecase.PickupReminderJob
}

func NewPickupHandler(
	usecase *usecase.PickupScheduleUsecase,
	reminderJob *usecase.PickupReminderJob,
) *PickupHandler {
	return &PickupHandler{
		usecase:     usecase,
		reminderJob: reminderJob,
	}
}

type createPickupRequest struct {
	UserID     int       `json:"user_id" example:"1"`
	PickupDate time.Time `json:"pickup_date" example:"2026-09-15T09:00:00Z"`
	WasteType  string    `json:"waste_type" example:"Organic"`
}

// CreatePickup godoc
// @Summary Create pickup schedule
// @Description Create a new waste pickup schedule
// @Tags Pickups
// @Accept json
// @Produce json
// @Param request body createPickupRequest true "Pickup schedule request"
// @Success 201 {object} usecase.PickupSchedule
// @Failure 400 {object} map[string]string
// @Router /api/v1/pickups [post]
func (h *PickupHandler) CreatePickup(c echo.Context) error {

	var request createPickupRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid request body",
			},
		)
	}

	schedule := &usecase.PickupSchedule{
		UserID:     request.UserID,
		PickupDate: request.PickupDate,
		WasteType:  request.WasteType,
	}

	if err := h.usecase.Create(
		c.Request().Context(),
		schedule,
	); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusCreated,
		schedule,
	)
}

// GetTomorrowPickup godoc
// @Summary Get tomorrow pickup schedules
// @Description Get all waste pickup schedules for tomorrow
// @Tags Pickups
// @Produce json
// @Success 200 {array} usecase.PickupSchedule
// @Failure 500 {object} map[string]string
// @Router /api/v1/pickups/tomorrow [get]
func (h *PickupHandler) GetTomorrowPickup(c echo.Context) error {

	schedules, err := h.usecase.GetTomorrow(
		c.Request().Context(),
	)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		schedules,
	)
}

// SendReminderTest godoc
// @Summary Test pickup reminder
// @Description Manually trigger the pickup reminder job for testing
// @Tags Pickups
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/pickups/reminder-test [post]
func (h *PickupHandler) SendReminderTest(c echo.Context) error {

	if h.reminderJob == nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": "pickup reminder job is not configured",
			},
		)
	}

	if err := h.reminderJob.SendTomorrowReminders(); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		map[string]string{
			"message": "pickup reminder sent successfully",
		},
	)
}
