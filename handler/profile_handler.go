package handler

import (
	"E-Meeting/internal/service"
	"E-Meeting/model"
	"database/sql"
	"net/http"

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
		c.String(http.StatusBadRequest, "error cuy")
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

func convertNullString(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}
