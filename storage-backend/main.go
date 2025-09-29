package main

import (
	"log"
	"os"

	"warehouse-go-backend/database"
	"warehouse-go-backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Подключение к БД
	database.ConnectDB()

	// 2. Инициализация Gin-роутера
	router := gin.Default()
	router.RedirectTrailingSlash = false
	
	// 3. Настройка CORS
	// Это критично для взаимодействия с фронтендом (Vue), работающим на другом порту (например, 8080)
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // В продакшене укажите конкретный домен Vue
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// 4. Определение маршрутов (Эндпоинты API)
	api := router.Group("/api/v1/inventory")
	{
		api.GET("/", handlers.GetItems)             // Получить все товары
		api.POST("/", handlers.CreateItem)          // Добавить новый товар
		api.DELETE("/:id", handlers.DeleteItem)     // Удалить товар по ID
		// api.PUT("/:id", handlers.UpdateItem)     // Обновить товар (логика не реализована в примере)
	}

	// 5. Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Порт по умолчанию
	}

	log.Printf("Сервер запущен на порту :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}