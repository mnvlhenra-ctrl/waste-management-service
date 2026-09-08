package delivery

import (
	"net/http"
	"waste-management-service/internal/domain"
	appMiddleware "waste-management-service/internal/middleware"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UUsecase domain.UserUsecase
}

func NewUserHandler(e *echo.Echo, us domain.UserUsecase) {
	handler := &UserHandler{UUsecase: us}

	// Routing Group API v1
	v1 := e.Group("/api/v1")

	users := v1.Group("/users")

	users.POST("/register", handler.Register) //nanti kebacanya jadi /api/v1/users/register
	users.POST("/login", handler.Login)

	// Endpoint yang butuh login
	users.GET("/profile", handler.GetProfile, appMiddleware.JWTMiddleware())

}

// register handler
func (h *UserHandler) Register(c echo.Context) error {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	ctx := c.Request().Context()
	if err := h.UUsecase.Register(ctx, &user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User registered successfully"})

}

// login handler
func (h *UserHandler) Login(c echo.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	ctx := c.Request().Context()
	token, err := h.UUsecase.Login(ctx, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})

}

// getprofile handler
func (h *UserHandler) GetProfile(c echo.Context) error {
	// 1. Ambil token dari context (disimpan oleh middleware echo-jwt)
	userToken, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token format"})
	}

	// 2. Ekstrak data claims
	claims := userToken.Claims.(jwt.MapClaims)

	// 3. Ambil ID (JWT secara default membaca angka sebagai float64)
	idFloat := claims["id"].(float64)
	userID := int(idFloat)

	// 4. Cari data profil di database
	ctx := c.Request().Context()
	profile, err := h.UUsecase.GetProfile(ctx, userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}

	// Jangan pernah mengirimkan password kembali ke frontend
	profile.Password = ""

	return c.JSON(http.StatusOK, profile)

}
