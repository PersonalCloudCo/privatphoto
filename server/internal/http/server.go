package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

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
		
		// Работа с файлами
		r.Get("/files", ListFilesHandler)
		r.Post("/files/upload", UploadHandler)
	})
	
	return r
}
