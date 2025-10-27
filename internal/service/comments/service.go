package comments

import (
	"context"

	"github.com/avraam311/comment-tree/internal/models"
)

type Repository interface {
	CreateComment(context.Context, *models.Comment) (int, error)
	GetAllComments(context.Context, int) ([]*models.CommentWithID, error)
	DeleteAllComments(context.Context, int) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
