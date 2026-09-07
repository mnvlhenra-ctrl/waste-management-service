package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token, ok := c.Get("user").(*jwt.Token)

		if !ok {
			return c.JSON(
				http.StatusUnauthorized,
				map[string]string{
					"message": "invalid token",
				},
			)
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			return c.JSON(
				http.StatusUnauthorized,
				map[string]string{
					"message": "invalid token claims",
				},
			)
		}

		role, ok := claims["role"].(string)

		if !ok || role != "ADMIN" {
			return c.JSON(
				http.StatusForbidden,
				map[string]string{
					"message": "admin access required",
				},
			)
		}

		return next(c)
	}
}
