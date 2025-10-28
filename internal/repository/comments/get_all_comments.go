package comments

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/avraam311/comment-tree/internal/models"
)

func (r *Repository) GetAllComments(ctx context.Context, parentID *uint) ([]*models.CommentWithID, error) {
	var query string
	var rows *sql.Rows
	var err error
	if parentID == nil {
		query = `
			SELECT id, text, parent_id
			FROM comment
			WHERE parent_id IS NULL
			ORDER BY created_at DESC;
		`
		rows, err = r.db.QueryContext(ctx, query)
		if err != nil {
			return nil, fmt.Errorf("repository/get_all_comments.go - failed to get comments - %w", err)
		}
		defer rows.Close()
	} else {
		query = `
			SELECT id, text, parent_id
			FROM comment
			WHERE parent_id = $1
			ORDER BY created_at DESC;
		`
		rows, err = r.db.QueryContext(ctx, query, parentID)
		if err != nil {
			return nil, fmt.Errorf("repository/get_all_comments.go - failed to get comments - %w", err)
		}
		defer rows.Close()
	}

	var coms []*models.CommentWithID
	for rows.Next() {
		c := models.CommentWithID{}
		var parentIDNull sql.NullInt64
		if err := rows.Scan(&c.ID, &c.Text, &parentIDNull); err != nil {
			return nil, fmt.Errorf("repository/get_all_comments.go - failed to scan comment - %w", err)
		}
		if parentIDNull.Valid {
			val := uint(parentIDNull.Int64)
			c.ParentID = &val
		} else {
			c.ParentID = nil
		}
		coms = append(coms, &c)
	}

	return coms, nil
}
