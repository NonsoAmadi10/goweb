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

func GetBook(c *gin.Context){
	//Get param id
	id := c.Param("id")
	var book models.Book

	if err := database.DB.Where("id = ? ", id).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context){
	var book models.Book
	//Get param id
	id := c.Param("id")

	

	if err := database.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	  }

	var input models.UpdateBookInput
	//validate input 
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
	}

	database.DB.Model(&book).Omit("ID").Updates(&input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context){
	// get param id 
	id := c.Param("id")

	var book models.Book

	if err := database.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	  }

	  database.DB.Delete(&book)
	  c.JSON(http.StatusOK, gin.H{"data": true})
}