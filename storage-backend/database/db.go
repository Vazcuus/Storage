package database

import (
	"fmt"
	"log"
	"os"

	"warehouse-go-backend/models" // Убедитесь, что путь к модели правильный

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Приоритет: использовать DATABASE_URL из окружения
	dsn := os.Getenv("DATABASE_URL")

	// Если DATABASE_URL не задан — собрать DSN из отдельных переменных
	if dsn == "" {
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		if dbHost == "" || dbPort == "" || dbUser == "" || dbPass == "" || dbName == "" {
			log.Fatal("Не указаны необходимые переменные окружения для подключения к БД")
		}

		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPass, dbName,
		)
	}

	// Подключение к БД
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Миграция (создание/обновление таблиц)
	err = db.AutoMigrate(&models.Item{})
	if err != nil {
		log.Fatalf("Ошибка миграции БД: %v", err)
	}

	DB = db
	fmt.Println("Успешное подключение к базе данных!")
}
