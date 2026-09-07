package main

import (
	"log"

	"waste-management-service/internal/config"
	transactionHTTP "waste-management-service/internal/transaction-microservice/delivery/http"
	transactionRepository "waste-management-service/internal/transaction-microservice/repository"
	transactionUsecase "waste-management-service/internal/transaction-microservice/usecase"
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

	// Transaction
	transactionRepo := transactionRepository.NewTransactionRepository(db)

	transactionUC := transactionUsecase.NewTransactionUsecase(
		transactionRepo,
	)

	transactionHandler := transactionHTTP.NewHandler(
		transactionUC,
	)

	// Waste
	wasteRepo := wasteRepository.NewWasteRepository(db)

	wasteUC := wasteUsecase.NewWasteUsecase(
		wasteRepo,
	)

	wasteHandler := wasteHTTP.NewHandler(
		wasteUC,
	)

	// Echo
	e := echo.New()

	transactionHTTP.RegisterRoutes(
		e,
		transactionHandler,
	)

	wasteHTTP.RegisterRoutes(
		e,
		wasteHandler,
	)

	log.Println("server running on :" + cfg.AppPort)

	if err := e.Start(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
