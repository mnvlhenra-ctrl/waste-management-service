package usecase

import (
	"context"
	"fmt"
	"time"

	"waste-management-service/pkg/email"
)

type PickupReminderService struct {
	emailService email.Service
}

func NewPickupReminderService(
	emailService email.Service,
) *PickupReminderService {
	return &PickupReminderService{
		emailService: emailService,
	}
}

func (s *PickupReminderService) SendReminder(
	ctx context.Context,
	userEmail string,
	username string,
	pickupDate time.Time,
	wasteType string,
) error {

	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<body style="font-family: Arial, sans-serif; line-height: 1.6;">

	<h2>Waste Pickup Reminder</h2>

	<p>Hello <strong>%s</strong>,</p>

	<p>
		This is a reminder that your waste will be collected tomorrow.
	</p>

	<table>
		<tr>
			<td><strong>Pickup Date</strong></td>
			<td>: %s</td>
		</tr>

		<tr>
			<td><strong>Waste Type</strong></td>
			<td>: %s</td>
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
`,
		username,
		pickupDate.Format("2006-01-02"),
		wasteType,
	)

	return s.emailService.Send(
		ctx,
		userEmail,
		"Waste Pickup Reminder - Tomorrow",
		body,
	)
}
