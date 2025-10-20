package service

import (
	"errors"
	"mime/multipart"
	"strings"

	"E-Meeting/internal/repository"
)

const (
	MaxFileSize   = 2048 * 2048            // 1MB
	SupportedType = "image/jpeg,image/png" // Hanya mendukung JPEG dan PNG
)

// UploadService mendefinisikan kontrak service upload
type UploadService interface {
	// UploadFile memproses validasi, menyimpan file, dan mengembalikan URL.
	UploadFile(fileHeader *multipart.FileHeader) (string, error)
}

// uploadService implementasi dari UploadService
type uploadService struct {
	repo repository.UploadRepository
}

// NewUploadService membuat instance UploadService baru
func NewUploadService(repo repository.UploadRepository) UploadService {
	return &uploadService{repo: repo}
}

// UploadFile menangani validasi dan penyimpanan file
func (s *uploadService) UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	// 1. VALIDASI UKURAN FILE
	if fileHeader.Size > MaxFileSize {
		return "", errors.New("file size too large/maximum file size is 1MB")
	}

	// 2. VALIDASI TIPE FILE
	contentType := fileHeader.Header.Get("Content-Type")
	// Kita hanya memeriksa header yang dikirim oleh klien.
	// Untuk keamanan penuh, Anda harus memeriksa magic number file (content-sniffing) juga.
	if !strings.Contains(SupportedType, contentType) {
		return "", errors.New("file type is not supported")
	}

	// 3. PANGGIL REPOSITORY UNTUK MENYIMPAN FILE
	imageURL, err := s.repo.SaveFile(fileHeader)
	if err != nil {
		// Log error di sini (jika menggunakan logger)
		return "", errors.New("gagal menyimpan file ke storage")
	}

	// 4. KEMBALIKAN HASIL
	return imageURL, nil
}
