package controllers

import (
	"net/http"
	"rakamin/app"
	"rakamin/database"
	"rakamin/models"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(context *gin.Context){

	conn, error := database.Connect()

	if error != nil{
		context.AbortWithStatusJSON(http.StatusBadGateway,gin.H{
			"status":404,
			"message":"Invalid credentials",
		})
		return
	}

	var newPhoto app.PhotoData
	if error := context.BindJSON(&newPhoto); error != nil{
		context.AbortWithStatusJSON(http.StatusBadGateway,gin.H{
			"status":404,
			"message":"Invalid credentials",
		})
		return
	}

	userData := context.MustGet("user").(models.User)

	insertPhoto := models.Photo{
		Title: newPhoto.Title,
		Caption: newPhoto.Caption,
		PhotoUrl: newPhoto.PhotoUrl,
		UserID: userData.ID,
	}

	conn.Create(&insertPhoto)

	context.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Successfully insert new photo",
	})
}

func GetPhoto(context *gin.Context){
	
	conn, error := database.Connect()

	if error != nil{
		context.AbortWithStatusJSON(http.StatusBadGateway,gin.H{
			"status":404,
			"message":"Invalid credentials",
		})	

		return
	}
	
	var photos []models.Photo 
	conn.First(&photos)

	context.IndentedJSON(http.StatusOK, gin.H{
		"result":photos,
		"status":200,
		"message":"Success",
	})
}

func UpdatePhoto(context *gin.Context){

	updateID := context.Param("id")

	conn, error := database.Connect()

	if error != nil{
		context.AbortWithStatusJSON(http.StatusBadGateway,gin.H{
			"status":404,
			"message":"Invalid credentials",
		})
		return
	}

	var newPhoto app.PhotoData
	if error := context.BindJSON(&newPhoto); error != nil{
		context.AbortWithStatusJSON(http.StatusBadGateway,gin.H{
			"status":404,
			"message":"Invalid credentials",
		})
		return
	}

	var photo models.Photo
	conn.Where("id = ?",updateID).First(&photo)

	photo.Title = newPhoto.Title
	photo.Caption = newPhoto.Caption
	photo.PhotoUrl = newPhoto.PhotoUrl

	conn.Save(&photo)

	context.IndentedJSON(http.StatusOK, gin.H{
		"status" : 200,
		"message": "Successfully update photo detail",
	})
}

func DeletePhoto(context *gin.Context){

	deleteID := context.Param("id")

	conn, error := database.Connect()

	if error != nil{
		context.AbortWithStatusJSON(http.StatusBadGateway,gin.H{
			"status":404,
			"message":"Invalid credentials",
		})
		return
	}

	var photo models.Photo
	conn.Where("id = ?",deleteID).First(&photo)

	conn.Delete(&photo)

	context.IndentedJSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Successfully delete photo",
	})
}

