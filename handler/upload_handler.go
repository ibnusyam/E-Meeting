package handler

import (
	"net/http"
	"strings"

	"E-Meeting/internal/service"

	"github.com/labstack/echo/v4"
)

type UploadHandler struct {
	uploadService service.UploadService
}

func NewUploadHandler(svc service.UploadService) *UploadHandler {
	return &UploadHandler{uploadService: svc}
}

func (h *UploadHandler) UploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "file field 'file' is required in form-data"})
	}

	imageURL, err := h.uploadService.UploadFile(file)
	if err != nil {
		message := err.Error()

		if strings.Contains(message, "file size too large") || strings.Contains(message, "file type is not supported") {
			// Bad Request: 400
			return c.JSON(http.StatusBadRequest, echo.Map{"message": message})
		}

		// Internal Server Error :500
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error: failed to process file"})
	}

	// RESPONSE SUKSES : 200
	return c.JSON(http.StatusOK, echo.Map{
		"message": "upload file success",
		"data": echo.Map{
			"imageURL": imageURL,
		},
	})
}
