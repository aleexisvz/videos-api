package handlers

import (
	"github.com/aleexisvz/videos-api/database"
	"github.com/aleexisvz/videos-api/models"
	"github.com/gin-gonic/gin"
)

// struct to handle RequestBody
type VideoRequestBody struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	Author string `json:"author"`
	Views  int    `json:"views"`
}

func VideoCreate(c *gin.Context) {
	// VideoRequestBody binding
	body := VideoRequestBody{}
	c.BindJSON(&body)

	// create a new video & insert into database
	video := models.Video{Title: body.Title, URL: body.URL, Author: body.Author}
	result := database.DB.Create(&video)

	// check for errors
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to insert"})
		return
	}

	// return video
	c.JSON(200, &video)
}

func VideoGet(c *gin.Context) {
	// create a new variable & get data from database
	var videos []models.Video
	database.DB.Find(&videos)

	// return videos
	c.JSON(200, &videos)
	return
}

func VideoGetById(c *gin.Context) {
	// get id
	id := c.Param("id")

	// create a new variable & get data from database
	var video models.Video
	database.DB.First(&video, id)

	VideoWatch(&video)

	// return the video
	c.JSON(200, &video)
	return
}

func VideoUpdate(c *gin.Context) {
	// get id
	id := c.Param("id")

	// create a new variable & get data from database
	var video models.Video
	database.DB.First(&video, id)

	// VideoRequestBody binding & create new variable to update
	body := VideoRequestBody{}
	c.BindJSON(&body)
	data := &models.Video{Title: body.Title, URL: body.URL, Author: body.Author}

	// update video to database
	result := database.DB.Model(&video).Updates(data)

	// check for errors
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update video"})
		return
	}

	// return the video
	c.JSON(200, &video)
}

func VideoDelete(c *gin.Context) {
	// get id
	id := c.Param("id")

	// create a new variable & get data from database
	var video models.Video
	database.DB.First(&video, id)

	// check if video exists
	if video.ID == 0 {
		c.JSON(404, gin.H{"Error": true, "message": "Video not found"})
		return
	}

	// delete video from database
	database.DB.Delete(&video)

	// return the video
	c.JSON(200, &video)
	return
}

func VideoWatch(video *models.Video) {
	// update video to database
	result := database.DB.Model(&video).Update("views", video.Views+1)

	// check for errors
	if result.Error != nil {
		return
	}
}
