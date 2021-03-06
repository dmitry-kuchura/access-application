package controllers

import (
	"strconv"
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

// Создание домена
func DomainCreate(c *gin.Context) {
	var data models.Domain
	if c.BindJSON(&data) == nil {
		id, err := models.CreateDomain(data.Name, data.Url, data.Description)

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

// Обновление информации по домену
func DomainUpdate(c *gin.Context) {
	var data models.Domain
	id, _ := strconv.Atoi(c.Param("id"))

	if c.BindJSON(&data) == nil {
		err := models.UpdateDomain(id, data.Name, data.Url, data.Description, data.Status)

		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"result":  "Domain was updated!",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"result":  "Not updated!",
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
	var data models.Domain
	if c.BindJSON(&data) == nil {
		_, err := models.DeleteDomain(data.ID)

		if err == nil {
			c.JSON(http.StatusNoContent, gin.H{
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

// Просмотр одного домена
func DomainView(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := models.GetDomain(id)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  data,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
		})
	}
}
