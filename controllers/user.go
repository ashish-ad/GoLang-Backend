package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/golang-backend/lib"
	"github.com/golang-backend/models"
)

type CreateUserInput struct {
	Name		string	`json:"name" binding:"required"`
	Surname string	`json:"surname" binding:"required"`
}

type UpdateUserInput struct {
	Name		string	`json:"name" binding:"required"`
	Surname string		`json:"surname" binding:"required"`
}

func CreateUser(c *gin.Context) { //Here c is context as req, res in express.js of NodeJs
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Name: input.Name, 
		Surname: input.Surname,
	}

	lib.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}


func FindUser(c *gin.Context) {
	var user models.User

	if err := lib.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := lib.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
	}

	var input UpdateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser := models.User{Name: input.Name, Surname: input.Surname}

	lib.DB.Model(&user).Updates(&updatedUser)
	c.JSON(http.StatusOK, gin.H{"Data": user})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := lib.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error":"record not found"})
		return
	}

	lib.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}