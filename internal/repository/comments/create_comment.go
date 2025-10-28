package comments

import (
	"context"
	"fmt"

	"github.com/avraam311/comment-tree/internal/models"
)

func (r *Repository) CreateComment(ctx context.Context, com *models.Comment) (uint, error) {
	query := `
		INSERT INTO comment (text, parent_id)
		VALUES ($1, $2)
		RETURNING id;
	`
	var id uint
	err := r.db.QueryRowContext(ctx, query, com.Text, com.ParentID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("repository/create_comment.go - failed to create comment - %w", err)
	}

	return id, nil
}
