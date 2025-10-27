package comments

import (
	"errors"

	"github.com/wb-go/wbf/dbpg"
)

var ErrCommentNotFound = errors.New("comment not found")

type Repository struct {
	db *dbpg.DB
}

func NewRepository(db *dbpg.DB) *Repository {
	return &Repository{
		db: db,
	}
}
