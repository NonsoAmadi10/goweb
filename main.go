package main

import (
	"github.com/gin-gonic/gin"
	"github.com/NonsoAmadi10/goweb/config"
	"github.com/NonsoAmadi10/goweb/models"
)


func main(){
	//initialize gin 

	r := gin.Default()

	// initialize database
	
	database.SetupDB(&models.Book{})

	r.Run()
}