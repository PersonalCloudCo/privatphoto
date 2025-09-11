package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// AuthMiddleware - заглушка для middleware аутентификации
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Пропускаем запрос без проверки (заглушка)
		next.ServeHTTP(w, r)
	})
}

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	// Добавляем базовые middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Регистрируем API маршруты
	r.Route("/api", func(r chi.Router) {
		// Аутентификация
		r.Post("/auth/login", LoginHandler)
		r.Post("/auth/verify-2fa", Verify2FAHandler)

		// Работа с файлами - защищаем middleware аутентификации
		r.Group(func(r chi.Router) {
			r.Use(AuthMiddleware)
			r.Get("/files", ListFilesHandler)
			r.Post("/files/upload", UploadHandler)
		})
	})

	return r
}