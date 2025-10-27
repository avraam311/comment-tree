package comments

import (
	"context"
	"fmt"

	"github.com/avraam311/comment-tree/internal/models"
)

func (s *Service) CreateComment(ctx context.Context, com *models.Comment) (int, error) {
	comID, err := s.repo.CreateComment(ctx, com)
	if err != nil {
		return 0, fmt.Errorf("service/create_comment.go - %w", err)
	}

	return comID, nil
}
