package main

import (
	"net/http"
	"database/sql"
	"./app"
	"./models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"fmt"
)

var db *sql.DB
var err error

func main() {
	if app.Config.Release == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/", Index)
	router.POST("/api/auth", Auth)
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	router.Run()
}

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
	})
}

func Auth(c *gin.Context) {
	var json models.User
	if c.BindJSON(&json) == nil {
		user := models.GetUser(json.Email)

		fmt.Print(user)

		if json.Email == "demo" && json.Password == "password" {
			c.JSON(http.StatusOK, gin.H{
				"status": "You are logged in",
				})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized",
			})
		}
	}
}
