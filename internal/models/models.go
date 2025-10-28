package models

type Comment struct {
	Text     string `json:"text" validate:"required"`
	ParentID *uint  `json:"parent_id,omitempty"`
}

type CommentWithID struct {
	ID       uint       `json:"id" validate:"required"`
	Text     string     `json:"text" validate:"required"`
	ParentID *uint      `json:"parent_id" validate:"required"`
	Children []*Comment `json:"children,omitempty"`
}
