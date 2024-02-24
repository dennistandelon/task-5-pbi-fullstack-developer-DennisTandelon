package database

import (
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
	"github.com/dennistandelon/task-5-pbi-fullstack-developer-DennisTandelon/models"
)

var (globalInstance *gorm.DB)

func Connect() (*gorm.DB, error) {
	
	if globalInstance != nil{
		return globalInstance, nil
	}
	
	err_load := godotenv.Load()

	if err_load != nil {
		println("Error to load the .env file")
		return nil, err_load
	}

	connection := os.Getenv("DB_CONNECTION")

	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		println("Error connecting to the database")
		return nil, err
	}

	db.AutoMigrate(&models.User{},&models.Photo{})

	globalInstance = db
	return db, nil
}