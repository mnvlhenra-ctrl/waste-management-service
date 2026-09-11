package http

import (
	"net/http"
	"time"

	"waste-management-service/pkg/email"

	"github.com/labstack/echo/v4"
)

type PickupReminderHandler struct {
	emailService email.Service
}

func NewPickupReminderHandler(
	emailService email.Service,
) *PickupReminderHandler {
	return &PickupReminderHandler{
		emailService: emailService,
	}
}

type pickupReminderRequest struct {
	Email      string `json:"email"`
	Username   string `json:"username"`
	PickupDate string `json:"pickup_date"`
	WasteType  string `json:"waste_type"`
}

func (h *PickupReminderHandler) TestReminder(
	c echo.Context,
) error {

	var request pickupReminderRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": "invalid request body",
			},
		)
	}

	if request.Email == "" {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": "email is required",
			},
		)
	}

	if request.Username == "" {
		request.Username = "Customer"
	}

	if request.PickupDate == "" {
		request.PickupDate = time.Now().
			AddDate(0, 0, 1).
			Format("2006-01-02")
	}

	if request.WasteType == "" {
		request.WasteType = "Mixed Waste"
	}

	pickupDate, err := time.Parse(
		"2006-01-02",
		request.PickupDate,
	)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": "invalid pickup_date format, use YYYY-MM-DD",
			},
		)
	}

	body := `
<!DOCTYPE html>
<html>
<body style="font-family: Arial, sans-serif; line-height: 1.6;">

	<h2>Waste Pickup Reminder</h2>

	<p>Hello <strong>` + request.Username + `</strong>,</p>

	<p>
		This is a reminder that your waste will be collected tomorrow.
	</p>

	<table>
		<tr>
			<td><strong>Pickup Date</strong></td>
			<td>: ` + pickupDate.Format("2006-01-02") + `</td>
		</tr>

		<tr>
			<td><strong>Waste Type</strong></td>
			<td>: ` + request.WasteType + `</td>
		</tr>
	</table>

	<p>
		Please make sure your waste is ready before the scheduled
		collection time.
	</p>

	<p>
		Thank you for using our waste management service.
	</p>

</body>
</html>
`

	err = h.emailService.Send(
		c.Request().Context(),
		request.Email,
		"Waste Pickup Reminder - Tomorrow",
		body,
	)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		map[string]string{
			"message": "pickup reminder email sent successfully",
		},
	)
}
