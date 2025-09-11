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

	// Создаем тестового пользователя, если его нет
	createTestUser()

	// Настраиваем и получаем HTTP-роутер
	r := http.SetupRouter()

	// Запускаем сервер
	log.Println("Server starting on :8080...")
	log.Fatal(nethttp.ListenAndServe(":8080", r))
}

func createTestUser() {
	// Проверяем, существует ли уже пользователь testuser
	user, err := storage.FindUserByLogin("testuser")
	if err != nil {
		log.Fatal("Failed to find user:", err)
	}
	if user == nil {
		// Создаем тестового пользователя
		_, err = storage.CreateUser("testuser", "testpassword", "test@example.com")
		if err != nil {
			log.Fatal("Failed to create test user:", err)
		}
		log.Println("Test user created: testuser / testpassword")
	} else {
		log.Println("Test user already exists: testuser")
	}
}