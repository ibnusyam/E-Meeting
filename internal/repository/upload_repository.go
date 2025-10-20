package repository

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type UploadRepository interface {
	SaveFile(fileHeader *multipart.FileHeader) (string, error)
}

type LocalDirectory struct {
	UploadDIR string
	BaseURL   string
}

func NewLocalDiskRepository(uploadDir, baseUrl string) *LocalDirectory {
	return &LocalDirectory{
		UploadDIR: uploadDir,
		BaseURL:   baseUrl,
	}
}

func (r *LocalDirectory) SaveFile(fileHeader *multipart.FileHeader) (string, error) {
	timestamp := time.Now().UnixNano()
	ext := filepath.Ext(fileHeader.Filename)

	filename := fmt.Sprintf("%d-%s%s", timestamp, strings.TrimSuffix(filepath.Base(fileHeader.Filename), ext), ext)

	src, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("gagal membuka file upload : %w", err)
	}
	defer src.Close()

	filePath := filepath.Join(r.UploadDIR, filename)

	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("gagal membuat file di disk : %w", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("gagal menyalin konten file : %w", err)
	}

	imageURL := r.BaseURL + "/" + filename
	return imageURL, nil

}
