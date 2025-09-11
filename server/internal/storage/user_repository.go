package storage

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// User представляет модель пользователя из БД
type User struct {
	ID           int
	Login        string
	PasswordHash string
	Email        sql.NullString
	TwoFAEnabled bool
	TwoFASecret  sql.NullString
	CreatedAt    sql.NullTime
}

// CreateUser создает нового пользователя с хешированным паролем
func CreateUser(login, password, email string) (int64, error) {
	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	// Вставляем пользователя в БД
	result, err := DB.Exec(
		"INSERT INTO users (login, password_hash, email) VALUES (?, ?, ?)",
		login, string(hashedPassword), email,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// FindUserByLogin ищет пользователя по логину
func FindUserByLogin(login string) (*User, error) {
	row := DB.QueryRow(
		"SELECT id, login, password_hash, email, two_fa_enabled, two_fa_secret, created_at FROM users WHERE login = ?",
		login,
	)

	var user User
	err := row.Scan(
		&user.ID,
		&user.Login,
		&user.PasswordHash,
		&user.Email,
		&user.TwoFAEnabled,
		&user.TwoFASecret,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Пользователь не найден
		}
		return nil, err
	}

	return &user, nil
}

// CheckPassword проверяет, соответствует ли пароль хешу
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}