package comments

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/avraam311/comment-tree/internal/api/handlers"
	"github.com/avraam311/comment-tree/internal/repository/comments"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) GetAllComments(c *ginext.Context) {
	parentIDStr := c.Query("parent")

	var parentIDpointer *uint
	if parentIDStr == "" {
		parentIDpointer = nil
	} else {
		parentID64, err := strconv.ParseUint(parentIDStr, 10, 64)
		if err != nil {
			zlog.Logger.Warn().Err(err).Msg("parent_id is not proper unsigned integer")
			handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid parent_id unsigned integer"))
			return
		}
		parentID := uint(parentID64)
		parentIDpointer = &parentID
	}

	coms, err := h.service.GetAllComments(c.Request.Context(), parentIDpointer)
	if err != nil {
		if errors.Is(err, comments.ErrCommentNotFound) {
			zlog.Logger.Warn().Err(err).Msg("comment not found")
			handlers.Fail(c.Writer, http.StatusNotFound, fmt.Errorf("comment not found"))
			return
		}

		zlog.Logger.Error().Err(err).Msg("failed to get comments")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.OK(c.Writer, coms)
}

func (h *Handler) DeleteAllComments(c *ginext.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		zlog.Logger.Warn().Err(err).Msg("id is not proper unsigned integer")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid id unsigned integer"))
		return
	}
	id := uint(id64)

	if err := h.service.DeleteAllComments(c.Request.Context(), id); err != nil {
		zlog.Logger.Warn().Err(err).Msg("failed to delete comments")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.OK(c.Writer, "comments deleted")
}
