package main

import (
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/golang-backend/controllers"
	"github.com/golang-backend/lib"
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
	router.POST("/create-post", controllers.CreateUser)
	router.GET("/find-user", controllers.FindUser)
	router.GET("/find-user/:id", controllers.FindUser)
	router.PATCH("/update-user/:id", controllers.UpdateUser)
	router.DELETE("/delete-user/:id", controllers.DeleteUser)

	lib.ConnectDatabase();

	// run the server
	router.Run(":"+PORT);
}