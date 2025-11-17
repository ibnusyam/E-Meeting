package handler

import (
	"E-Meeting/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	Service *service.AuthService
}

func NewLoginHandler(service *service.AuthService) *LoginHandler {
	return &LoginHandler{Service: service}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login godoc
// @Summary Login
// @Description Authenticate user and return access and refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login Request"
// @Success 200 {object} map[string]interface{} "Login Success"
// @Failure 400 {object} map[string]string "Login Failed"
// @Router /login [post]
func (h *LoginHandler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Login Failed",
		})
	}

	accessToken, refreshToken, err := h.Service.Login(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Login Failed",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login Success",
		"data": map[string]string{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	})
}
