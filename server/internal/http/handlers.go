package http

import (
	"encoding/json"
	"net/http"

	"github.com/PersonalCloudCo/privatphoto/server/internal/storage"
)

// LoginRequest представляет структуру запроса на логин
type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// LoginHandler обрабатывает запрос на аутентификацию
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Ищем пользователя в БД
	user, err := storage.FindUserByLogin(req.Login)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if user == nil || !user.CheckPassword(req.Password) {
		http.Error(w, "Invalid login or password", http.StatusUnauthorized)
		return
	}

	// Пока заглушка: всегда возвращаем requires2fa = false
	response := map[string]interface{}{
		"success":     true,
		"requires2fa": false,
		"token":       "dummy_jwt_token_for_now", // Временный токен
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Verify2FAHandler обрабатывает проверку 2FA кода
func Verify2FAHandler(w http.ResponseWriter, r *http.Request) {
	// Заглушка - всегда возвращаем успех
	response := map[string]interface{}{
		"success":      true,
		"sessionToken": "dummy_session_token_for_now",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UploadHandler обрабатывает загрузку файлов
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Заглушка - всегда возвращаем успех
	response := map[string]interface{}{
		"success": true,
		"fileId":  1,
	}

	json.NewEncoder(w).Encode(response)
}

// ListFilesHandler возвращает список файлов пользователя
func ListFilesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Заглушка - возвращаем пустой массив файлов
	files := []map[string]interface{}{}

	json.NewEncoder(w).Encode(files)
}