package handlers

import (
	"net/http"
	"warehouse-go-backend/database"
	"warehouse-go-backend/models"

	"github.com/gin-gonic/gin"
)

// GetItems godoc
// @Summary Получить список всех товаров
func GetItems(c *gin.Context) {
	var items []models.Item
	// Поиск всех записей в таблице items
	if result := database.DB.Find(&items); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить список товаров"})
		return
	}
	// Отправляем массив товаров в формате JSON
	c.JSON(http.StatusOK, items)
}

// CreateItem godoc
// @Summary Создать новый товар
func CreateItem(c *gin.Context) {
	var input models.NewItemRequest
	// Привязка JSON тела запроса к структуре
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	newItem := models.Item{
		Name: input.Name,
		Quantity: input.Quantity,
	}

	// Создание записи в БД
	if result := database.DB.Create(&newItem); result.Error != nil {
		// Обработка ошибки дубликата имени
		c.JSON(http.StatusConflict, gin.H{"error": "Товар с таким именем уже существует."})
		return
	}

	c.JSON(http.StatusCreated, newItem)
}

// DeleteItem godoc
// @Summary Удалить товар по ID
func DeleteItem(c *gin.Context) {
	// Получаем ID из URL-параметра
	id := c.Param("id")
	
	// Удаление записи (GORM использует "мягкое" удаление, если присутствует gorm.Model)
	if result := database.DB.Delete(&models.Item{}, id); result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Товар успешно удален"})
}

// Note: Для полной реализации также необходимы GetItemByID и UpdateItem