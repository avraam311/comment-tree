package comments

import (
	"context"
	"fmt"

	"github.com/avraam311/comment-tree/internal/models"
)

func (s *Service) GetAllComments(ctx context.Context, parentID int) ([]*models.CommentWithID, error) {
	coms, err := s.repo.GetAllComments(ctx, parentID)
	if err != nil {
		return nil, fmt.Errorf("service/get_all_comments - %w", err)
	}

	return coms, nil
}
