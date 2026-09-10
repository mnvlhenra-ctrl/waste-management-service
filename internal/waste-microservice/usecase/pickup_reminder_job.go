package usecase

import (
	"context"
	"log"

	"waste-management-service/internal/domain"
)

type PickupReminderJob struct {
	pickupRepo  PickupScheduleRepository
	userRepo    domain.UserRepository
	reminderSvc *PickupReminderService
}

func NewPickupReminderJob(
	pickupRepo PickupScheduleRepository,
	userRepo domain.UserRepository,
	reminderSvc *PickupReminderService,
) *PickupReminderJob {
	return &PickupReminderJob{
		pickupRepo:  pickupRepo,
		userRepo:    userRepo,
		reminderSvc: reminderSvc,
	}
}

func (j *PickupReminderJob) SendTomorrowReminders() error {

	ctx := context.Background()

	schedules, err := j.pickupRepo.GetTomorrow(ctx)
	if err != nil {
		return err
	}

	for _, schedule := range schedules {

		user, err := j.userRepo.GetByID(
			ctx,
			schedule.UserID,
		)

		if err != nil {
			log.Printf(
				"failed to get user %d: %v",
				schedule.UserID,
				err,
			)
			continue
		}

		err = j.reminderSvc.SendReminder(
			ctx,
			user.Email,
			user.Email,
			schedule.PickupDate,
			schedule.WasteType,
		)

		if err != nil {
			log.Printf(
				"failed to send pickup reminder to %s: %v",
				user.Email,
				err,
			)
			continue
		}

		log.Printf(
			"pickup reminder sent to %s for %s",
			user.Email,
			schedule.PickupDate.Format("2006-01-02"),
		)
	}

	return nil
}
