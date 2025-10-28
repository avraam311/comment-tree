package comments

import (
	"context"

	"github.com/avraam311/comment-tree/internal/models"
	"github.com/go-playground/validator/v10"
)

type Service interface {
	CreateComment(context.Context, *models.Comment) (uint, error)
	GetAllComments(context.Context, *uint) ([]*models.CommentWithID, error)
	DeleteAllComments(context.Context, uint) error
}

type Handler struct {
	service   Service
	validator *validator.Validate
}

func NewHandler(service Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}
