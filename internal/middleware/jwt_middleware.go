package middleware

import (
	"net/http"
	"strings"

	"E-Meeting/internal/utils"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Missing Authorization Header",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		_, err := utils.ValidateAccessToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Invalid or expired token",
			})
		}

		return next(c)
	}
}
