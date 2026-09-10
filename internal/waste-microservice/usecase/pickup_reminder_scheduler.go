package usecase

import (
	"context"
	"log"
	"time"
)

type PickupReminderScheduler struct {
	job *PickupReminderJob
}

func NewPickupReminderScheduler(
	job *PickupReminderJob,
) *PickupReminderScheduler {
	return &PickupReminderScheduler{
		job: job,
	}
}

func (s *PickupReminderScheduler) Start(ctx context.Context) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println("failed to load timezone:", err)
		return
	}

	for {
		now := time.Now().In(loc)

		// Target setiap hari jam 09:00 WIB.
		next := time.Date(
			now.Year(),
			now.Month(),
			now.Day(),
			9,
			0,
			0,
			0,
			loc,
		)

		// Kalau hari ini jam 09:00 sudah lewat,
		// jadwalkan untuk besok.
		if !next.After(now) {
			next = next.AddDate(0, 0, 1)
		}

		wait := time.Until(next)

		log.Printf(
			"next pickup reminder scheduled at %s",
			next.Format("2006-01-02 15:04:05 MST"),
		)

		timer := time.NewTimer(wait)

		select {
		case <-ctx.Done():
			timer.Stop()
			log.Println("pickup reminder scheduler stopped")
			return

		case <-timer.C:
			if err := s.job.SendTomorrowReminders(); err != nil {
				log.Println("pickup reminder error:", err)
			}
		}
	}
}
