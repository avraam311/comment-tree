package comments

import (
	"context"
	"fmt"

	"github.com/avraam311/comment-tree/internal/models"
)

func (r *Repository) GetAllComments(ctx context.Context, parentID int) ([]*models.CommentWithID, error) {
	query := `
		SELECT id, text, parent_id
		FROM comment
		WHERE id = $1 OR parent_id = $1;
	`

	rows, err := r.db.QueryContext(ctx, query, parentID)
	if err != nil {
		return nil, fmt.Errorf("repository/get_all_comments.go - failed to get comments - %w", err)
	}
	defer rows.Close()

	coms := make([]*models.CommentWithID, 0)
	for rows.Next() {
		c := &models.CommentWithID{}
		if err := rows.Scan(&c.ID, &c.Text, &c.ParentID); err != nil {
			return nil, fmt.Errorf("repository/get_all_comments.go - failed to scan comments rows - %w", err)
		}

		coms = append(coms, c)
	}

	return coms, nil
}
