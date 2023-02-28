package handlers

import (
	"net/http"
	"nocode/src/db"
	"nocode/src/models"

	"github.com/gin-gonic/gin"
)

var forms []models.Form

func CreateForm(c *gin.Context) {
	var form models.Form
	err := c.BindJSON(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	forms = append(forms, form)
	c.JSON(http.StatusCreated, form)
}

func GetForms(c *gin.Context) {
	forms, err := db.QueryAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server Error"})
		return
	}
	c.JSON(http.StatusOK, forms)
}
