package main

import (
	"log"
	nethttp "net/http" // Переименовываем стандартный пакет http чтобы избежать конфликта

	"github.com/PersonalCloudCo/privatphoto/server/internal/http" // Импорт нашего внутреннего пакета
	"github.com/PersonalCloudCo/privatphoto/server/internal/storage"
)

func main() {
	// Инициализируем БД
	storage.InitDB()
	// Настраиваем и получаем HTTP-роутер
	r := http.SetupRouter()

	// Запускаем сервер
	log.Println("Server starting on :8080...")
	log.Fatal(nethttp.ListenAndServe(":8080", r))
}