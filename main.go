package main

import (
	"fmt"
	"time"

	"net/http"

	"./app"
	"./models"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
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

	config := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},
		MaxAge: 12 * time.Hour,
	})

	router := gin.Default()
	router.GET("/", Index)
	router.POST("/api/auth", Auth)

	router.DELETE("/api/user-delete", UserDelete)
	router.POST("/api/user-create", UserCreate)
	router.POST("/api/user-change-status", UserChangeStatus)
	router.GET("/api/user-list", UsersGetList)

	router.POST("/api/domain-create", DomainCreate)
	router.GET("/ws", func(c *gin.Context) {
		WebSocketsHandler(c.Writer, c.Request)
	})

	router.Use(config)

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

// Регистрация пользователей
func UserCreate(c *gin.Context) {
	var data models.User
	if c.BindJSON(&data) == nil {
		_, err := models.CreateUser(data.Email, data.Password, data.Name)

		if err == nil {
			c.JSON(http.StatusCreated, gin.H{
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

// Плучение полного списка пользователей
func UsersGetList(c *gin.Context) {
	list, err := models.AllUsers()

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"users":   list,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
		})
	}
}

// Создание домена
func DomainCreate(c *gin.Context) {
	var data models.Domains
	if c.BindJSON(&data) == nil {
		id, err := models.CreateDomain(data.Name, data.Url)

		if err == nil {
			c.JSON(http.StatusCreated, gin.H{
				"success": true,
				"domain":  id,
				"result":  "Domain was registered!",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"result":  "Not created",
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
