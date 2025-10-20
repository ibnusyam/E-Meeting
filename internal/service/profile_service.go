package service

import (
	"E-Meeting/internal/repository"
	"E-Meeting/model"
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type ProfileService struct {
	Repo *repository.ProfileRepository
}

func NewProfileService(repo *repository.ProfileRepository) *ProfileService {
	return &ProfileService{Repo: repo}
}

func (s *ProfileService) GetUserProfileByID(id string) (model.User, error) {
	profile, err := s.Repo.GetUserProfileByID(id)
	if err != nil {
		return profile, err
	}

	return profile, nil
}

// fungsi untuk mengubah profile user
var ErrUserNotFound = repository.ErrUserNotFound // Memperoleh error 404 dari repo

// UpdateUser memperbarui user
func (s *ProfileService) UpdateUser(ctx context.Context, id int, req model.UserUpdateRequest) (model.UserResponse, error) {
	// ...

	if req.Password != nil && *req.Password != "" {

		// Hashing password baru
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return model.UserResponse{}, fmt.Errorf("failed to hash password: %w", err)
		}

		// Simpan hash string kembali ke pointer req.Password
		// sehingga Repository akan menyimpan hash, bukan plaintext
		hashedString := string(hashedPassword)
		req.Password = &hashedString
	}

	// 3. Panggil Repository
	user, err := s.Repo.UpdateUserPartial(ctx, id, req)
	if err != nil {
		// Ini mencari ErrEmailTaken di package 'repository'
		if errors.Is(err, repository.ErrEmailTaken) {
			return model.UserResponse{}, errors.New("email is already taken") // 400
		}
		return model.UserResponse{}, fmt.Errorf("service: failed to update user: %w", err)
	}

	return user, nil
}
