package controllers

import (
	"net/http"

	"github.com/NonsoAmadi10/goweb/config"
	"github.com/NonsoAmadi10/goweb/models"
	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var books []models.Book

	database.DB.Find(books)

	c.JSON(http.StatusOK, gin.H{ "data": books })
}