package routes

import (
	"nocode/src/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterFormRoutes(router *gin.Engine) {
	router.POST("/forms", handlers.CreateForm)
	router.GET("forms", handlers.GetForms)
}
