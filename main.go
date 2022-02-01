package main

import (
	"github.com/gin-gonic/gin"
	"github.com/NonsoAmadi10/goweb/config"
	"github.com/NonsoAmadi10/goweb/models"
	"github.com/NonsoAmadi10/goweb/controllers"
)


func main(){
	//initialize gin 

	r := gin.Default()

	// initialize database
	
	database.SetupDB(&models.Book{})

	r.GET("/books", controllers.FindBooks)

	r.Run()
}