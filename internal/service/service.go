package service

import (
	"fmt"
	"live/internal/repository"
	"live/models"

	"github.com/sirupsen/logrus"
)

type Service struct {
	repository *repository.Repository
	queue      chan *models.Comment
}

func NewService(repository *repository.Repository) *Service {
	service := &Service{
		repository: repository,
		queue:      make(chan *models.Comment, 100),
	}

	for i := 0; i < 3; i++ {
		go service.worker(i)
	}

	return service
}

func (s *Service) AddComment(comment models.Comment) (models.Comment, error) {
	comment, err := s.repository.CreateNewComment(comment)
	if err != nil {
		return models.Comment{}, err
	}

	select {
	case s.queue <- &comment:
		logrus.Info("Comment added to queue: ", comment.ID)
	default:
		logrus.Error("queue full")
		return models.Comment{}, fmt.Errorf("queue full")
	}

	return comment, nil
}

func (s *Service) worker(id int) {
	for comment := range s.queue {
		s.repository.UpdateCommentStatus(comment.ID)
		fmt.Printf("Worker %d finished comment %d\n", id, comment.ID)
	}
}
