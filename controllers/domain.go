package controllers

import (
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

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

// Список доменов
func DomainList(c *gin.Context) {
	list, count, err := models.AllDomains(c.Param("page"))

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"pages":   count,
			"domains": list,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
		})
	}
}

// Удаление домена
func DomainDelete(c *gin.Context) {
	var data models.Domains
	if c.BindJSON(&data) == nil {
		_, err := models.DeleteDomain(data.ID)

		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"result":  "Domain was deleted!",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"result":  "Not deleted",
			})
		}
	}
}