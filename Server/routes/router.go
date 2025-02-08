package router

import (
	"Server/controllers"
	"Server/server"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(server *server.Server) {
	server.Router.GET("/search", controllers.SearchMovie)
	server.Router.POST("/post", func(c *gin.Context) {
		controllers.PostMovie(c, server.WatchlistCollection)
	})
}
