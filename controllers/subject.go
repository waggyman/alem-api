package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/waggyman/alem-api/models"
)

func CreateSubject(c *gin.Context) {
	var newSubject models.Subject
	if err := c.BindJSON(&newSubject); err != nil {
		return
	}
	models.StoreSubject(newSubject)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "New subject added"})
}
