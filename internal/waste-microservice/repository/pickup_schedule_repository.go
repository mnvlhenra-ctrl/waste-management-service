package repository

import (
	"context"
	"time"

	"waste-management-service/internal/waste-microservice/usecase"

	"gorm.io/gorm"
)

type pickupScheduleModel struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	UserID     int       `gorm:"not null"`
	PickupDate time.Time `gorm:"not null"`
	WasteType  string    `gorm:"not null"`
}

func (pickupScheduleModel) TableName() string {
	return "pickup_schedules"
}

type pickupScheduleRepository struct {
	db *gorm.DB
}

func NewPickupScheduleRepository(
	db *gorm.DB,
) usecase.PickupScheduleRepository {

	// Buat tabel pickup_schedules jika belum ada.
	if err := db.AutoMigrate(&pickupScheduleModel{}); err != nil {
		panic("failed to migrate pickup_schedules table: " + err.Error())
	}

	return &pickupScheduleRepository{
		db: db,
	}
}

func (r *pickupScheduleRepository) Create(
	ctx context.Context,
	schedule *usecase.PickupSchedule,
) error {

	model := &pickupScheduleModel{
		UserID:     schedule.UserID,
		PickupDate: schedule.PickupDate,
		WasteType:  schedule.WasteType,
	}

	if err := r.db.WithContext(ctx).
		Create(model).Error; err != nil {
		return err
	}

	schedule.ID = model.ID

	return nil
}

func (r *pickupScheduleRepository) GetTomorrow(
	ctx context.Context,
) ([]usecase.PickupSchedule, error) {

	var models []pickupScheduleModel

	now := time.Now()

	startOfTomorrow := time.Date(
		now.Year(),
		now.Month(),
		now.Day()+1,
		0,
		0,
		0,
		0,
		now.Location(),
	)

	startOfDayAfterTomorrow := startOfTomorrow.AddDate(0, 0, 1)

	err := r.db.WithContext(ctx).
		Where(
			"pickup_date >= ? AND pickup_date < ?",
			startOfTomorrow,
			startOfDayAfterTomorrow,
		).
		Order("pickup_date ASC").
		Find(&models).Error

	if err != nil {
		return nil, err
	}

	schedules := make(
		[]usecase.PickupSchedule,
		0,
		len(models),
	)

	for _, model := range models {
		schedules = append(
			schedules,
			usecase.PickupSchedule{
				ID:         model.ID,
				UserID:     model.UserID,
				PickupDate: model.PickupDate,
				WasteType:  model.WasteType,
			},
		)
	}

	return schedules, nil
}
