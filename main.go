package main

import (
	"fmt"

	"net/http"

	"./app"
	"./models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var WebSocketsRefresher = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	switch app.Config.Release {
	case "DebugMode":
		gin.SetMode(gin.DebugMode)
		break
	case "TestMode":
		gin.SetMode(gin.TestMode)
		break
	case "ReleaseMode":
		gin.SetMode(gin.ReleaseMode)
		break
	}

	router := gin.Default()
	router.GET("/", Index)
	router.DELETE("/api/user-delete", UserDelete)
	router.POST("/api/user-create", UserCreate)
	router.POST("/api/user-change-status", UserChangeStatus)
	router.POST("/api/auth", Auth)
	router.GET("/ws", func(c *gin.Context) {
		WebSocketsHandler(c.Writer, c.Request)
	})

	if app.Config.ServerPort == "" {
		router.Run()
	} else {
		router.Run(app.Config.ServerPort)
	}
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

// Удаление пользователей
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
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"result":  "Not deleted",
			})
		}
	}

}

// Изменение статуса пользователя
func UserChangeStatus(c *gin.Context) {
	var data models.User
	if c.BindJSON(&data) == nil {
		user, err := models.ChangeStatusUser(data.ID)

		if err == false {
			c.JSON(http.StatusOK, gin.H{
				"success":        true,
				"current_status": user.Status,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"result":  "Not change",
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

func WebSocketsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := WebSocketsRefresher.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}
