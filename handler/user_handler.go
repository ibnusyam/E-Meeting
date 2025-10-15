package handler

import (
	"E-Meeting/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service}
}

type RegisterRequest struct {
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

// Register godoc
// @Summary      Register user baru
// @Description  Register user baru
// @Tags         Register
// @Accept       json
// @Produce      json
// @Param        request  body  RegisterRequest  true  "Register Request"
// @Success      200  {object}  map[string]interface{}  "Register Success"
// @Failure      400  {object}  map[string]interface{}  "Bad Request"
// @Failure      500  {object}  map[string]interface{}  "Internal Server Error"
// @Router       /register [post]
func (h *UserHandler) Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Register Failed"})
	}

	err := h.service.Register(req.Email, req.Username, req.Password, req.ConfirmPassword)
	if err != nil {
		if err.Error() == "passwords do not match" || err.Error() == "email already registered" {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Internal Server Error"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Register Success",
		"data":    nil,
	})
}
