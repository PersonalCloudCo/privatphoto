package http

import (
	"encoding/json"
	"net/http"
)

// LoginHandler обрабатывает запрос на аутентификацию
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Заглушка - всегда возвращаем успех и флаг requires2fa = false
	response := map[string]interface{}{
		"success":     true,
		"requires2fa": false,
		"token":       "dummy_jwt_token_for_now",
	}
	
	json.NewEncoder(w).Encode(response)
}

// Verify2FAHandler обрабатывает проверку 2FA кода
func Verify2FAHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Заглушка - всегда возвращаем успех
	response := map[string]interface{}{
		"success":      true,
		"sessionToken": "dummy_session_token_for_now",
	}
	
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
