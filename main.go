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
	router.POST("/api/auth", Auth)
	router.Run()
}

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
	})
}

func UserCreate(c *gin.Context) {
	var data models.User

	user, err := models.CreateUser(data.Email, data.Password, data.Name)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"user_id": user,
			"result":  "You account was registered!",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"result":  "Not registered",
		})
	}

}

func Auth(c *gin.Context) {
	var json models.User
	if c.BindJSON(&json) == nil {
		user := models.GetUser(json.Email, json.Password)

		fmt.Println(user.ID)
		fmt.Println(user.Token)

		if json.Email == "demo" && json.Password == "password" {
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
