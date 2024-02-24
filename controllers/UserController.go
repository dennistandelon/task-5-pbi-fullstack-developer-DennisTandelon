package controllers

import (
	"net/http"
	"rakamin/app"
	"rakamin/database"
	"rakamin/helpers"
	"rakamin/models"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	conn, error := database.Connect()

	if error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Error while connection to the database",
		})
		return
	}

	var newUser app.UserData
	if error := context.BindJSON(&newUser); error != nil {
		context.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status":  404,
			"message": "Invalid credentials",
		})
		return
	}

	insertUser := models.User{
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: helpers.EncryptPassword(newUser.Password),
	}

	conn.Create(&insertUser)

	context.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Successfully registered",
	})
}

func Login(context *gin.Context) {
	conn, error := database.Connect()

	if error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Error while connection to the database",
		})

		return
	}

	var user models.User

	email := context.Query("email")
	password := context.Query("password")

	err := conn.Where("email = ?", email).First(&user).Error

	if err != nil || !helpers.CheckPassword(password, user.Password) {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  404,
			"message": "Invalid credentials",
		})
		return
	}

	token, err := helpers.GenerateToken(user)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Error while generating token",
		})
		return
	}

	context.SetCookie("Authorization", token, 3600, "", "", true, true)

	context.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Successfully logged in",
	})
}

func Logout(context *gin.Context) {

	_, err := context.Cookie("Authorization")

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"message": "Unauthorized",
		})
		return
	}

	context.SetCookie("Authorization", "", -1, "", "", true, true)

	context.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Successfully Logout",
	})
}

func UpdateUser(context *gin.Context) {

	conn, error := database.Connect()

	if error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Error while connection to the database",
		})
		return
	}

	updateID := context.Param("id")

	var newUser app.UserData
	if error := context.BindJSON(&newUser); error != nil {
		context.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status":  404,
			"message": "Invalid credentials",
		})
		return
	}
	
	userData := context.MustGet("user").(models.User)

	var user models.User
	conn.Where("id = ?", updateID).First(&user)


	if user.ID == userData.ID{
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
			"status":401,
			"message":"You dont have access to delete this photo",
		})
		return
	}

	user.Username = newUser.Username
	user.Email = newUser.Email
	user.Password = helpers.EncryptPassword(newUser.Password)

	context.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Successfully update user",
	})
}

func DeleteUser(context *gin.Context) {

	conn, error := database.Connect()

	if error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Error while connection to the database",
		})
		return
	}

	deleteID := context.Param("id")

	var user models.User
	conn.Where("id = ?", deleteID).First(&user)

	userData := context.MustGet("user").(models.User)

	if user.ID == userData.ID{
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
			"status":401,
			"message":"You dont have access to delete this photo",
		})
		return
	}

	conn.Delete(&user)

	context.IndentedJSON(http.StatusOK, gin.H{
		"message": "delete user",
	})
}
