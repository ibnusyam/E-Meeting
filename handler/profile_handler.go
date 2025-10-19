package handler

import (
	"E-Meeting/internal/service"
	"E-Meeting/model"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ProfileHandler struct {
	Service *service.ProfileService
}

func NewProfileHandler(service *service.ProfileService) *ProfileHandler {
	return &ProfileHandler{Service: service}
}

// handler/profile_handler.go

// GetAllSnacks godoc
// @Summary      Mendapatkan data profile
// @Description  Mengambil data profile berdasarkan id
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.ProfileUser
// @Router       /profile/:id [get]
func (h *ProfileHandler) GetUserProfileByID(c echo.Context) error {
	id := c.Param("id")
	profile, err := h.Service.GetUserProfileByID(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "user not found")
	}
	profileDTO := model.ProfileUser{
		ID:             profile.ID,
		Role:           profile.Role,
		Username:       profile.Username,
		Email:          profile.Email,
		PhoneNumber:    profile.PhoneNumber,
		Language:       profile.Language,
		Status:         profile.Status,
		ProfilePicture: convertNullString(profile.ProfilePicture),
		CreatedAt:      profile.CreatedAt,
		UpdatedAt:      profile.UpdatedAt,
	}

	return c.JSON(http.StatusOK, profileDTO)
}

func (h *ProfileHandler) UpdateUserHandler(c echo.Context) error {
	idParam := c.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid user ID format"})
	}

	var req model.UserUpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request body"})
	}

	// Panggil Service
	updatedUser, err := h.Service.UpdateUser(c.Request().Context(), userID, req)

	if err != nil {
		log.Printf("User Update Error [ID: %d]: %v", userID, err)

		if errors.Is(err, service.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "user not found"}) // 404
		}
		// Asumsi error lain dari service adalah Bad Request/Internal
		if strings.Contains(err.Error(), "email is already taken") {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "email already taken"}) // 400
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "internal server error"}) // 500
	}

	// Respons Sukses
	return c.JSON(http.StatusOK, echo.Map{
		"message": "user updated successfully",
		"data":    updatedUser,
	})
}

func convertNullString(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}
