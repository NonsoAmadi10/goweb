package controllers

import (
	"net/http"

	"github.com/NonsoAmadi10/goweb/config"
	"github.com/NonsoAmadi10/goweb/models"
	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var books []models.Book

	database.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{ "data": books })
}

func CreateBook(c *gin.Context){

	//validate input
	var input models.CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create book 
	book := models.Book{Title: input.Title, Author: input.Author }

	database.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{" data": book})
}