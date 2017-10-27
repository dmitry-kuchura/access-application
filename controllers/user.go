package controllers

import (
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

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
func UserList(c *gin.Context) {
	list, count, err := models.AllUsers(c.Param("page"))

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"pages":   count,
			"users":   list,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
		})
	}
}
