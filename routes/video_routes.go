package routes

import (
	"github.com/aleexisvz/videos-api/handlers"
	"github.com/gin-gonic/gin"
)

func VideoRouter(router *gin.Engine) {
	routes := router.Group("/videos")

	routes.POST("", handlers.VideoCreate)
	routes.GET("", handlers.VideoGet)
	routes.GET("/:id", handlers.VideoGetById)
	routes.PUT("/:id", handlers.VideoUpdate)
	routes.DELETE("/:id", handlers.VideoDelete)
}
