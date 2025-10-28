package comments

import (
	"context"

	"github.com/avraam311/comment-tree/internal/models"
)

type Repository interface {
	CreateComment(context.Context, *models.Comment) (uint, error)
	GetAllComments(context.Context, *uint) ([]*models.CommentWithID, error)
	DeleteAllComments(context.Context, uint) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
