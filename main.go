package main;

import (
	"net/http"
	"github.com/golang-backend/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default(); // router with default middleware installed
	// index route

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	});

	models.ConnectDatabase();

	// run the server
	router.Run();
}