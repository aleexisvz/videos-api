package main

import (
	"github.com/aleexisvz/videos-api/database"
	"github.com/aleexisvz/videos-api/models"
)

// init invokes the ConnectToDB from database package
func init() {
	database.ConnectToDB()
}

// main invokes the AutoMigrate from database package using Video model
func main() {
	database.DB.AutoMigrate(&models.Video{})
}
