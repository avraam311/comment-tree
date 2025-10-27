package comments

import (
	"context"
	"fmt"
)

func (r *Repository) DeleteAllComments(ctx context.Context, id int) error {
	query := `
		DELETE
		FROM comment
		WHERE id = $1 OR parent_id = $1;
    `

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("repository/delete_all_comments.go - failed to delete comments - %w", err)
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return ErrCommentNotFound
	}

	return nil
}
