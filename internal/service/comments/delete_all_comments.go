package comments

import (
	"context"
	"fmt"
)

func (s *Service) DeleteAllComments(ctx context.Context, id int) error {
	err := s.repo.DeleteAllComments(ctx, id)
	if err != nil {
		return fmt.Errorf("service/get_analytics - %w", err)
	}

	return nil
}
