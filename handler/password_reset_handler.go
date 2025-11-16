package handler

import (
	"E-Meeting/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PasswordResetHandler struct {
	Service *service.PasswordResetService
}

func NewPasswordResetHandler(service *service.PasswordResetService) *PasswordResetHandler {
	return &PasswordResetHandler{Service: service}
}

// POST /password/reset
// Request password reset godoc
// @Summary      Request password reset
// @Description  Mengirim permintaan untuk mereset password dengan email
// @Tags         Password Reset
// @Accept       json
// @Produce      json
// @Param        request body map[string]string true "Email Request"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /password/reset [post]
func (h *PasswordResetHandler) ResetRequest(c echo.Context) error {
	var req struct {
		Email string `json:"email"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
	}

	token, err := h.Service.GenerateResetToken(req.Email)
	if err != nil {
		if err.Error() == "email not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "email not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Update Password success",
		"data": map[string]string{
			"token": token,
		},
	})
}

// PUT /password/reset/:id
// Reset password by id godoc
// @Summary      Reset password by id
// @Description  Mereset password menggunakan ID user
// @Tags         Password Reset
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        request body map[string]string true "Password Reset Request"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /password/reset/{id} [put]
func (h *PasswordResetHandler) ResetPassword(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		NewPassword     string `json:"new_password"`
		ConfirmPassword string `json:"confirm_password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
	}

	err := h.Service.UpdatePassword(id, req.NewPassword, req.ConfirmPassword)
	if err != nil {
		if err.Error() == "password confirmation is not match" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "password confirmation is not match"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Update Password success"})
}
