package repository

import (
	"E-Meeting/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

type ProfileRepository struct {
	DB *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{DB: db}
}

func (repo *ProfileRepository) GetUserProfileByID(id string) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := "SELECT id, role, username, email, password, phone_number, language, status_user, profile_picture, created_at, updated_at FROM USERS where id=" + id

	rows := repo.DB.QueryRowContext(ctx, query)

	var user model.User

	err := rows.Scan(&user.ID, &user.Role, &user.Username, &user.Email, &user.Password, &user.PhoneNumber, &user.Language, &user.Status, &user.ProfilePicture, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

// fungsi untuk mengubah userprofile
// variabel error
var ErrUserNotFound = errors.New("user not found")
var ErrEmailTaken = errors.New("email already taken")

func (r *ProfileRepository) UpdateUserPartial(ctx context.Context, id int, req model.UserUpdateRequest) (model.UserResponse, error) {
	// Membangun query secara dinamis
	sets := []string{}
	args := []interface{}{}
	argCounter := 1

	if req.Username != nil {
		sets = append(sets, fmt.Sprintf("username = $%d", argCounter)) // Mengubah kolom 'name'
		args = append(args, *req.Username)
		argCounter++
	}

	// 3. EMAIL
	if req.Email != nil {
		sets = append(sets, fmt.Sprintf("email = $%d", argCounter))
		args = append(args, *req.Email)
		argCounter++
	}

	// 4. PASSWORD (Diasumsikan sudah di-hash di Service Layer)
	if req.Password != nil {
		// Di sini Anda mengasumsikan kolom di DB adalah 'password_hash' atau 'password'
		sets = append(sets, fmt.Sprintf("password = $%d", argCounter))
		args = append(args, *req.Password) // Ini adalah nilai hash
		argCounter++
	}

	// 5. PROFILE PICTURE / Image URL
	if req.ProfilePicture != nil {
		sets = append(sets, fmt.Sprintf("profile_picture = $%d", argCounter))
		args = append(args, *req.ProfilePicture)
		argCounter++
	}

	// 6. PHONE NUMBER
	if req.PhoneNumber != nil {
		sets = append(sets, fmt.Sprintf("phone_number = $%d", argCounter))
		args = append(args, *req.PhoneNumber)
		argCounter++
	}

	// 7. LANGUAGE
	if req.Language != nil {
		sets = append(sets, fmt.Sprintf("language = $%d", argCounter))
		args = append(args, *req.Language)
		argCounter++
	}

	// 8. ROLE
	if req.Role != nil {
		sets = append(sets, fmt.Sprintf("role = $%d", argCounter))
		args = append(args, *req.Role)
		argCounter++
	}

	// 9. STATUS
	if req.Status != nil {
		sets = append(sets, fmt.Sprintf("status = $%d", argCounter))
		args = append(args, *req.Status)
		argCounter++
	}

	if len(sets) == 0 {
		return model.UserResponse{}, errors.New("no fields to update")
	}

	// Tambahkan updated_at (tidak menambah argCounter karena NOW() adalah fungsi DB)
	sets = append(sets, "updated_at = NOW()")

	// Tambahkan ID ke argumen terakhir
	args = append(args, id)

	// 🔥 argCounter saat ini adalah posisi placeholder ID pengguna
	query := fmt.Sprintf(`
        UPDATE users 
        SET %s 
        WHERE id = $%d 
        RETURNING id, username, created_at, email, password, profile_picture, language, role, status_user, updated_at, phone_number
    `,
		strings.Join(sets, ", "),
		argCounter)

	// ...

	// Asumsi: Kita menggunakan variabel temporer untuk kolom NULL, dan UserResponse
	// memiliki field string biasa (yang kemudian diisi di logika if valid/else).
	var user model.UserResponse
	var nullProfilePicture sql.NullString // Asumsi profile_picture Nullable

	err := r.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,            // 1. id
		&user.Username,      // 2. username
		&user.CreatedAt,     // 3. created_at
		&user.Email,         // 4. email
		&user.Password,      // 5. password (hash)
		&nullProfilePicture, // 6. profile_picture (ke Nullable)
		&user.Language,      // 7. language
		&user.Role,          // 8. role
		&user.Status,        // 9. status_user (Pastikan field di model Anda bernama Status)
		&user.UpdatedAt,     // 10. updated_at
		&user.PhoneNumber,   // 11. phone_number
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// user tidak ditemukan
			return model.UserResponse{}, ErrUserNotFound
		}

		// Handle error unik (misalnya UNIQUE constraint violation)
		if strings.Contains(err.Error(), "unique constraint") {
			return model.UserResponse{}, ErrEmailTaken
		}

		return model.UserResponse{}, fmt.Errorf("repo: failed to execute update and scan: %w", err)
	}

	// Mengembalikan data user yang sudah di-update
	return user, nil
}
