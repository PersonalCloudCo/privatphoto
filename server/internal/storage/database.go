package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Драйвер SQLite
)

var DB *sql.DB

func InitDB() {
	var err error
	// Открываем соединение с SQLite БД (файл будет создан автоматически)
	DB, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	// Проверяем соединение
	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Database connection established")
}