package handler

import (
	"net/http"

	"live/internal/service"

	"live/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		service: *services,
	}
}

func (h *Handler) InitRout() *gin.Engine {
	router := gin.Default()

	router.POST("/comment", h.AddComment)

	return router
}

func (h *Handler) AddComment(c *gin.Context) {
	var comment models.Comment
	err := c.Bind(&comment)
	if err != nil {
		logrus.WithError(err).Error("failed parse request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed parse request body"})
		return
	}
	comment.Status = "on moderated"
	err = h.service.AddComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "add comment failed"})
		return
	}

	c.JSON(http.StatusOK, comment)
}
