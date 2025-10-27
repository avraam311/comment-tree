package comments

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/avraam311/comment-tree/internal/api/handlers"
	"github.com/avraam311/comment-tree/internal/models"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) CreateComment(c *ginext.Context) {
	var com models.Comment

	if err := json.NewDecoder(c.Request.Body).Decode(&com); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to decode request body")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid request body: %s", err.Error()))
		return
	}

	if err := h.validator.Struct(com); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to validate request body")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("validation error: %s", err.Error()))
		return
	}

	parentID := strconv.FormatUint(uint64(com.ParentID), 10)
	if com.ParentID < 1 {
		zlog.Logger.Error().Str("parent_id", parentID).Msg("parent_id less than 1")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("parent_id must be greater than 1: %s", parentID))
		return
	}

	id, err := h.service.CreateComment(c.Request.Context(), &com)
	if err != nil {
		zlog.Logger.Error().Err(err).Interface("comment", com).Msg("failed to create comment")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.Created(c.Writer, id)
}
