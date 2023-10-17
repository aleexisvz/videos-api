package main

import (
	"net/http"

	"github.com/aleexisvz/videos-api/database"
	"github.com/aleexisvz/videos-api/routes"
	"github.com/gin-gonic/gin"
)

// init invokes the ConnectToDB from database package
func init() {
	database.ConnectToDB()
}

func main() {
	// create a new gin router
	r := gin.Default()

	// video routes
	routes.VideoRouter(r)

	// default route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	// start server
	r.Run(":8080")
}
