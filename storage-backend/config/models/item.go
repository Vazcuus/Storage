package models

import "gorm.io/gorm"

// Item представляет запись о товаре на складе
type Item struct {
	gorm.Model           // Включает поля ID, CreatedAt, UpdatedAt, DeletedAt
	Name      string     `json:"name" gorm:"unique;not null"`
	Quantity  int        `json:"quantity" gorm:"not null;check:quantity >= 0"`
}

// NewItemRequest используется для валидации входных данных при создании
type NewItemRequest struct {
	Name      string `json:"name" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,gt=0"`
}

// UpdateItemRequest используется для обновления данных
type UpdateItemRequest struct {
	Name      string `json:"name"`
	Quantity  int    `json:"quantity" binding:"required,gte=0"`
}