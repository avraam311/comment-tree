package comments

import (
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	validator *validator.Validate
}

func NewHandler(validator *validator.Validate) *Handler {
	return &Handler{
		validator: validator,
	}
}
