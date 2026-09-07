package main

import (
	"context"
	"log"

	"waste-management-service/internal/config"
	transactionHTTP "waste-management-service/internal/transaction-microservice/delivery/http"
	transactionRepository "waste-management-service/internal/transaction-microservice/repository"
	transactionUsecase "waste-management-service/internal/transaction-microservice/usecase"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.LoadConfig()

	dbURL := "postgres://" +
		cfg.DBUser + ":" +
		cfg.DBPassword + "@" +
		cfg.DBHost + ":" +
		cfg.DBPort + "/" +
		cfg.DBName

	db, err := pgxpool.New(
		context.Background(),
		dbURL,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	transactionRepo := transactionRepository.NewPostgresRepository(db)

	transactionUC := transactionUsecase.NewTransactionUsecase(
		transactionRepo,
	)

	transactionHandler := transactionHTTP.NewHandler(
		transactionUC,
	)

	e := echo.New()

	transactionHTTP.RegisterRoutes(
		e,
		transactionHandler,
	)

	log.Println("server running on :" + cfg.AppPort)

	if err := e.Start(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
