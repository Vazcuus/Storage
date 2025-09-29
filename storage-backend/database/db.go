package database

import (
	"fmt"
	"log"
	"os"

	"warehouse-go-backend/models" // Замените на ваш фактический путь

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// DSN (Data Source Name) для подключения к PostgreSQL
	// Рекомендуется хранить в переменных среды (env variables)
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// DSN по умолчанию (замените на ваши данные)
		dsn = "host=localhost user=lev password=123 dbname=postgres port=5440 sslmode=disable TimeZone=Asia/Tomsk"
	}
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Автоматическое создание/обновление таблицы (миграция)
	err = db.AutoMigrate(&models.Item{})
	if err != nil {
		log.Fatalf("Ошибка миграции БД: %v", err)
	}

	DB = db
	fmt.Println("Успешное подключение к базе данных!")
}