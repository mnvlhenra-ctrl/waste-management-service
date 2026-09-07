package http

import (
	"net/http"
	"strconv"

	"waste-management-service/internal/transaction-microservice/usecase"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase *usecase.TransactionUsecase
}

func NewHandler(
	usecase *usecase.TransactionUsecase,
) *Handler {

	return &Handler{
		usecase: usecase,
	}
}

type topUpRequest struct {
	UserID int     `json:"user_id"`
	Amount float64 `json:"amount"`
}

type paymentRequest struct {
	UserID    int     `json:"user_id"`
	InvoiceID int     `json:"invoice_id"`
	Amount    float64 `json:"amount"`
}

// CreateTopUp godoc
// @Summary Create top up transaction
// @Description Create a new top up transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Param request body topUpRequest true "Top up request"
// @Success 201 {object} domain.Transaction
// @Failure 400 {object} map[string]string
// @Router /transactions/topup [post]
func (h *Handler) CreateTopUp(
	c echo.Context,
) error {

	var request topUpRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid request body",
			},
		)
	}

	transaction, err := h.usecase.CreateTopUp(
		c.Request().Context(),
		request.UserID,
		request.Amount,
	)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusCreated,
		transaction,
	)
}

// CreatePayment godoc
// @Summary Create payment transaction
// @Description Create a new payment transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Param request body paymentRequest true "Payment request"
// @Success 201 {object} domain.Transaction
// @Failure 400 {object} map[string]string
// @Router /transactions/pay [post]
func (h *Handler) CreatePayment(
	c echo.Context,
) error {

	var request paymentRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid request body",
			},
		)
	}

	transaction, err := h.usecase.CreatePayment(
		c.Request().Context(),
		request.UserID,
		request.InvoiceID,
		request.Amount,
	)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusCreated,
		transaction,
	)
}

// GetAll godoc
// @Summary Get user transactions
// @Description Get all transactions for a user
// @Tags Transactions
// @Produce json
// @Param user_id query int true "User ID"
// @Success 200 {array} domain.Transaction
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions [get]
func (h *Handler) GetAll(
	c echo.Context,
) error {

	userID, err := strconv.Atoi(
		c.QueryParam("user_id"),
	)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid user id",
			},
		)
	}

	transactions, err := h.usecase.GetAll(
		c.Request().Context(),
		userID,
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
		transactions,
	)
}

// GetByID godoc
// @Summary Get transaction by ID
// @Description Get transaction by ID
// @Tags Transactions
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} domain.Transaction
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /transactions/{id} [get]
func (h *Handler) GetByID(
	c echo.Context,
) error {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid transaction id",
			},
		)
	}

	transaction, err := h.usecase.GetByID(
		c.Request().Context(),
		id,
	)

	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		transaction,
	)
}
