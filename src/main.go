package main

import (
	"nocode/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RegisterFormRoutes(router)
	router.Run()
}
