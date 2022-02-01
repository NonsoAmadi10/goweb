package database

import (
	"github.com/NonsoAmadi10/goweb/helpers"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	
)

var DB *gorm.DB

func SetupDB(model...interface{}){
	dbString := helpers.GetEnv("DATABASE_URL")
	database, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})

	if err != nil {
		panic("Failed to Connect to Database")
	}

	database.AutoMigrate(model...)

	DB = database
}