package main

import (
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/golang-backend/controllers"
	"github.com/golang-backend/models"
)

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	router := gin.Default(); // router with default middleware installed
	// index route


	//Test router
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	});

	//ALl the routers are defined here
	router.POST("/create-post", controllers.CreatePost)
	

	models.ConnectDatabase();

	// run the server
	router.Run(":"+PORT);
}