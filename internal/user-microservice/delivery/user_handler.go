package delivery

import (
	"context"
	"net/http"

	"waste-management-service/internal/domain"
	appMiddleware "waste-management-service/internal/middleware"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type invoiceService interface {
	GetByUserID(ctx context.Context, userID int) ([]domain.Invoice, error)
}

type UserHandler struct {
	UUsecase       domain.UserUsecase
	InvoiceUsecase invoiceService
}

func NewUserHandler(
	e *echo.Echo,
	us domain.UserUsecase,
	invoiceUC invoiceService,
) {
	handler := &UserHandler{
		UUsecase:       us,
		InvoiceUsecase: invoiceUC,
	}

	v1 := e.Group("/api/v1")

	users := v1.Group("/users")

	users.POST("/register", handler.Register)
	users.POST("/login", handler.Login)

	// Endpoint yang butuh login
	users.GET(
		"/profile",
		handler.GetProfile,
		appMiddleware.JWTMiddleware(),
	)

	// Get invoice milik user yang sedang login
	users.GET("/invoices", handler.GetMyInvoices, appMiddleware.JWTMiddleware())
}

type registerRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	HouseNumber string `json:"house_number"`
}

func (h *UserHandler) Register(c echo.Context) error {

	var request registerRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": "invalid request body",
			},
		)
	}

	user := domain.User{
		Email:    request.Email,
		Password: request.Password,
	}

	ctx := c.Request().Context()

	if err := h.UUsecase.Register(
		ctx,
		&user,
		request.HouseNumber,
	); err != nil {

		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusCreated,
		map[string]string{
			"message": "User registered successfully",
		},
	)
}

func (h *UserHandler) Login(c echo.Context) error {

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": err.Error(),
			},
		)
	}

	token, err := h.UUsecase.Login(
		c.Request().Context(),
		req.Email,
		req.Password,
	)

	if err != nil {
		return c.JSON(
			http.StatusUnauthorized,
			map[string]string{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		map[string]string{
			"token": token,
		},
	)
}

func (h *UserHandler) GetProfile(c echo.Context) error {

	userToken, ok := c.Get("user").(*jwt.Token)

	if !ok {
		return c.JSON(
			http.StatusUnauthorized,
			map[string]string{
				"message": "Invalid token format",
			},
		)
	}

	claims, ok := userToken.Claims.(jwt.MapClaims)

	if !ok {
		return c.JSON(
			http.StatusUnauthorized,
			map[string]string{
				"message": "Invalid token claims",
			},
		)
	}

	idFloat, ok := claims["id"].(float64)

	if !ok {
		return c.JSON(
			http.StatusUnauthorized,
			map[string]string{
				"message": "Invalid user id",
			},
		)
	}

	userID := int(idFloat)

	profile, err := h.UUsecase.GetProfile(
		c.Request().Context(),
		userID,
	)

	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			map[string]string{
				"message": "User not found",
			},
		)
	}

	// Jangan kirim password
	profile.Password = ""

	return c.JSON(
		http.StatusOK,
		profile,
	)
}

// GetMyInvoices mengambil semua invoice milik user yang sedang login
func (h *UserHandler) GetMyInvoices(c echo.Context) error {

	userToken, ok := c.Get("user").(*jwt.Token)

	if !ok {
		return c.JSON(
			http.StatusUnauthorized,
			map[string]string{
				"message": "Invalid token format",
			},
		)
	}

	claims, ok := userToken.Claims.(jwt.MapClaims)

	if !ok {
		return c.JSON(
			http.StatusUnauthorized,
			map[string]string{
				"message": "Invalid token claims",
			},
		)
	}

	idFloat, ok := claims["id"].(float64)

	if !ok {
		return c.JSON(
			http.StatusUnauthorized,
			map[string]string{
				"message": "Invalid user id",
			},
		)
	}

	userID := int(idFloat)

	invoices, err := h.InvoiceUsecase.GetByUserID(
		c.Request().Context(),
		userID,
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
		invoices,
	)
}
