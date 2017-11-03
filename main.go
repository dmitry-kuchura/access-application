package main

import (
	"fmt"
	"time"

	"net/http"

	"./app"
	"./models"
	"./controllers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/gorilla/websocket"
	"log"
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
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		MaxAge:       12 * time.Hour,
	})

	router := gin.New()

	// Global middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/", Index)
	router.NoRoute(PageNotFound)

	router.POST("/api/auth", Auth)

	user := router.Group("/api/user")
	user.Use(AuthRequired())
	{
		user.DELETE("delete", controllers.UserDelete)
		user.POST("create", controllers.UserCreate)
		user.PUT("update/:id", controllers.UserUpdate)
		user.POST("change-status", controllers.UserChangeStatus)
		user.GET("list/:page", controllers.UserList)
	}

	domains := router.Group("/api/domain")
	domains.Use(AuthRequired())
	{
		domains.POST("create", controllers.DomainCreate)
		domains.PUT("update/:id", controllers.DomainUpdate)
		domains.GET("view/:id", controllers.DomainView)
		domains.GET("list/:page", controllers.DomainList)
		domains.DELETE("delete", controllers.DomainDelete)
	}

	ftp := router.Group("/api/ftp")
	ftp.Use(AuthRequired())
	{
		ftp.POST("create", controllers.FtpCreate)
		ftp.DELETE("delete", controllers.FtpDelete)
	}

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

// Middleware авторизация
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before authRequired")
		token := c.Request.Header.Get("Auth-Token")

		log.Println("Token: " + token)
		_, err := models.FindUserByToken(token)

		// No token No Party
		if err == true {
			log.Println("no auth")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   http.StatusUnauthorized,
				"message": "invalid Auth-Token!",
			})
		}
		log.Println("after authRequired")
	}
}

// Кастомная Not Found (404)
func PageNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error":   404,
		"message": "Not found",
	})
}

// Главная страница
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

// Хандлер для WebSocket'ов
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
