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

	// Создаем таблицу пользователей, если она не существует
	createTable()
	log.Println("Database connection established")
}

func createTable() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		login TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		email TEXT,
		two_fa_enabled BOOLEAN DEFAULT FALSE,
		two_fa_secret TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}
}