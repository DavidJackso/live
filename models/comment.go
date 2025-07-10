package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `json:"title" binding:"required"`
	Status  string `json:"status"`
}
