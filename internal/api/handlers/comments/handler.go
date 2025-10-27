package comments

import (
	"context"

	"github.com/avraam311/comment-tree/internal/models"
	"github.com/go-playground/validator/v10"
)

type Service interface {
	CreateComment(context.Context, *models.Comment) (int, error)
	GetAllComments(context.Context, int) ([]*models.CommentWithID, error)
	DeleteAllComments(context.Context, int) error
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
