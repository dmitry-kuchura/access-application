package main

import (
	"./app"
	"./models"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	if app.Config.Release == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.GET("/", Index)
	router.POST("/api/user-create", UserCreate)
	router.POST("/api/user-delete", UserDelete)
	router.POST("/api/auth", Auth)
	router.Run()
}

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
	})
}

// Регистрация пользователей
func UserCreate(c *gin.Context) {
	var data models.User
	if c.BindJSON(&data) == nil {
		_, err := models.CreateUser(data.Email, data.Password, data.Name)

		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"result":  "You account was registered!",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"result":  "Not registered",
			})
		}
	}

}

func UserDelete(c *gin.Context) {
	var data models.User
	if c.BindJSON(&data) == nil {
		_, err := models.DeleteUser(data.ID)

		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"result":  "You account was deleted!",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"result":  "Not deleted",
			})
		}
	}

}

// Авторизация пользователей
func Auth(c *gin.Context) {
	var json models.User
	if c.BindJSON(&json) == nil {
		user, err := models.GetUser(json.Email, json.Password)

		if err != true {
			c.JSON(http.StatusOK, gin.H{
				"status": "You are logged in",
				"token":  user.Token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized",
			})
		}
	}
}
