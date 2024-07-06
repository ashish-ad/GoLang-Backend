package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/golang-backend/models"
)

type CreateUserInput struct {
	Name		string	`json:"title" binding:"required"`
	Surname string	`json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Name: input.Name, 
		Surname: input.Surname,
	}

	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}