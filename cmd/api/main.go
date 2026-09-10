package main

import (
	"context"
	"log"

	"waste-management-service/internal/config"

	transactionHTTP "waste-management-service/internal/transaction-microservice/delivery/http"
	transactionRepository "waste-management-service/internal/transaction-microservice/repository"
	transactionUsecase "waste-management-service/internal/transaction-microservice/usecase"

	"waste-management-service/pkg/email"

	userHTTP "waste-management-service/internal/user-microservice/delivery"
	userRepository "waste-management-service/internal/user-microservice/repository"
	userUsecase "waste-management-service/internal/user-microservice/usecase"

	repository "waste-management-service/internal/utils/repository"

	wasteHTTP "waste-management-service/internal/waste-microservice/delivery/http"
	wasteRepository "waste-management-service/internal/waste-microservice/repository"
	wasteUsecase "waste-management-service/internal/waste-microservice/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	dbURL := "host=" + cfg.DBHost +
		" user=" + cfg.DBUser +
		" password=" + cfg.DBPassword +
		" dbname=" + cfg.DBName +
		" port=" + cfg.DBPort +
		" sslmode=disable"

	db, err := gorm.Open(
		postgres.Open(dbURL),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal(err)
	}

	// =========================
	// Echo
	// =========================
	e := echo.New()

	// =========================
	// Email Service - Resend
	// =========================
	emailSvc := email.NewResendService()

	// =========================
	// User
	// =========================
	uRepo := userRepository.NewUserRepository(db)

	// =========================
	// House & Invoice
	// =========================
	houseRepo := repository.NewHouseRepository(db)
	invoiceRepo := repository.NewInvoiceRepository(db)

	// =========================
	// User Usecase
	// =========================
	uUC := userUsecase.NewUserUsecase(
		uRepo,
		houseRepo,
		invoiceRepo,
		emailSvc,
		cfg.JWTSecret,
	)

	invoiceUC := userUsecase.NewInvoiceUsecase(
		invoiceRepo,
	)

	userHTTP.NewUserHandler(
		e,
		uUC,
		invoiceUC,
	)

	// =========================
	// Transaction
	// =========================
	transactionRepo := transactionRepository.NewTransactionRepository(db)

	transactionUC := transactionUsecase.NewTransactionUsecase(
		transactionRepo,
		uRepo,
		emailSvc,
	)

	transactionHandler := transactionHTTP.NewHandler(
		transactionUC,
	)

	transactionHTTP.RegisterRoutes(
		e,
		transactionHandler,
	)

	// =========================
	// Waste
	// =========================
	wasteRepo := wasteRepository.NewWasteRepository(db)

	wasteUC := wasteUsecase.NewWasteUsecase(
		wasteRepo,
	)

	wasteHandler := wasteHTTP.NewHandler(
		wasteUC,
	)

	wasteHTTP.RegisterRoutes(
		e,
		wasteHandler,
	)

	// =========================
	// Pickup Schedule
	// =========================
	pickupRepo := wasteRepository.NewPickupScheduleRepository(db)

	pickupUC := wasteUsecase.NewPickupScheduleUsecase(
		pickupRepo,
	)

	// =========================
	// Pickup Reminder
	// =========================
	pickupReminderSvc := wasteUsecase.NewPickupReminderService(
		emailSvc,
	)

	pickupReminderJob := wasteUsecase.NewPickupReminderJob(
		pickupRepo,
		uRepo,
		pickupReminderSvc,
	)

	// =========================
	// Pickup Handler
	// =========================
	pickupHandler := wasteHTTP.NewPickupHandler(
		pickupUC,
		pickupReminderJob,
	)

	wasteHTTP.RegisterPickupRoutes(
		e,
		pickupHandler,
	)

	// =========================
	// Pickup Reminder Scheduler
	// =========================
	pickupReminderScheduler := wasteUsecase.NewPickupReminderScheduler(
		pickupReminderJob,
	)

	go pickupReminderScheduler.Start(context.Background())

	// =========================
	// Start Server
	// =========================
	log.Println("server running on :" + cfg.AppPort)

	if err := e.Start(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
