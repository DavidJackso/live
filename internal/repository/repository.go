package repository

import (
	"fmt"
	"live/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateNewComment(comment models.Comment) (models.Comment, error) {
	comment.Status = "on_moderation"
	result := r.db.Create(&comment)
	if result.Error != nil {
		return models.Comment{}, result.Error
	}
	fmt.Println(comment)
	return comment, nil
}

func (r *Repository) UpdateCommentStatus(id uint) error {
	var comment models.Comment
	result := r.db.Where("id = ?", id).First(&comment)

	if result.Error != nil {
		return result.Error
	}
	comment.Status = "moderated"
	if err := r.db.Save(&comment).Error; err != nil {
		return fmt.Errorf("failed update comment status: %w", err)
	}

	return nil
}
