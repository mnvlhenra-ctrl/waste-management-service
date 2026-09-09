package main

import (
	"log"

	"waste-management-service/internal/config"

	transactionHTTP "waste-management-service/internal/transaction-microservice/delivery/http"
	transactionRepository "waste-management-service/internal/transaction-microservice/repository"
	transactionUsecase "waste-management-service/internal/transaction-microservice/usecase"

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

	// Echo
	e := echo.New()

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
		cfg.JWTSecret,
	)

	userHTTP.NewUserHandler(
		e,
		uUC,
	)

	// =========================
	// Transaction
	// =========================
	transactionRepo := transactionRepository.NewTransactionRepository(db)

	transactionUC := transactionUsecase.NewTransactionUsecase(
		transactionRepo,
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
	// Start Server
	// =========================
	log.Println("server running on :" + cfg.AppPort)

	if err := e.Start(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
