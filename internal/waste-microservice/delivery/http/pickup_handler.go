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
	UserID     int       `json:"user_id"`
	PickupDate time.Time `json:"pickup_date"`
	WasteType  string    `json:"waste_type"`
}

// CreatePickup membuat jadwal pengambilan sampah.
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

// GetTomorrowPickup mengambil jadwal pickup untuk besok.
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

// SendReminderTest menjalankan pickup reminder secara manual untuk testing.
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
