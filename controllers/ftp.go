package controllers

import (
	"net/http"
	"../models"
	"github.com/gin-gonic/gin"
)

// Создание домена
func FtpCreate(c *gin.Context) {
	var data models.Ftp
	if c.BindJSON(&data) == nil {
		id, err := models.CreateFtp(data.DomainID, data.Hostname, data.Username, data.Password)

		if err == nil {
			c.JSON(http.StatusCreated, gin.H{
				"success": true,
				"domain":  id,
				"result":  "Created!",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"result":  "Not created",
			})
		}
	}
}

// Удаление домена
func FtpDelete(c *gin.Context) {
	var data models.Domain
	if c.BindJSON(&data) == nil {
		_, err := models.DeleteFtp(data.ID)

		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"result":  "Was deleted!",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"result":  "Not deleted",
			})
		}
	}
}
