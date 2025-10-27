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

	if parentIDStr == "" {
		zlog.Logger.Error().Msg("no parameter parent recieved")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("parameter parent is required"))
		return
	}

	parentID, err := strconv.Atoi(parentIDStr)
	if err != nil {
		zlog.Logger.Warn().Err(err).Msg("parent_id is not proper integer")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid parent_id integer"))
		return
	}

	coms, err := h.service.GetAllComments(c.Request.Context(), parentID)
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
	id, err := strconv.Atoi(idStr)
	if err != nil {
		zlog.Logger.Warn().Err(err).Msg("id is not proper integer")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid id integer"))
		return
	}

	if id < 0 {
		zlog.Logger.Warn().Err(err).Msg("negative id")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("id must be >= 0"))
		return
	}

	if err := h.service.DeleteAllComments(c.Request.Context(), id); err != nil {
		zlog.Logger.Warn().Err(err).Msg("failed to delete comments")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.OK(c.Writer, "comments deleted")
}
