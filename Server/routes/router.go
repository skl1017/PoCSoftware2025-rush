package router

import (
	"Server/controllers"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/search", controllers.SearchMovie)
	r.POST("/post", controllers.PostMovie)
}
