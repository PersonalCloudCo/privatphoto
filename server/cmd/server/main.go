package main

import (
	"log"
	"net/http"
	"privatphoto/server/internal/http" // Импорт нашего внутреннего пакета
	"privatphoto/server/internal/storage"
)

func main() {
	// Инициализируем БД
	storage.InitDB()
	// Настраиваем и получаем HTTP-роутер
	r := http.SetupRouter()

	// Запускаем сервер
	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}