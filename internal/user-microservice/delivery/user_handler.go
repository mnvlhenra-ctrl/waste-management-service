package delivery

import (
	"waste-management-service/internal/domain"

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
	users.GET("/profile", handler.GetProfile)
}

// register handler
func (h *UserHandler) Register(c echo.Context) error {

}

// login handler
func (h *UserHandler) Login(c echo.Context) error {

}

// getprofile handler
func (h *UserHandler) GetProfile(c echo.Context) error {

}
