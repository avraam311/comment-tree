package server

import (
	"net/http"

	"github.com/wb-go/wbf/ginext"

	"github.com/avraam311/comment-tree/internal/api/handlers/comments"
	"github.com/avraam311/comment-tree/internal/middlewares"
)



func NewRouter(ginMode string, handlerComs *comments.Handler) *ginext.Engine {
	e := ginext.New(ginMode)

	e.Use(middlewares.CORSMiddleware())
	e.Use(ginext.Logger())
	e.Use(ginext.Recovery())

	api := e.Group("/comment-tree/api")
	{
		api.POST("/comments", handlerComs.CreateComment)
		api.GET("comments", handlerComs.GetAllComments)
		api.DELETE("/comments/:id", handlerComs.DeleteAllComments)
	}

	return e
}

func NewServer(addr string, router *ginext.Engine) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
