package main

import (
	"fmt"
	"nocode/src/db"
	"nocode/src/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	db.InitDb()
	router := gin.Default()

	routes.RegisterFormRoutes(router)
	router.Run()
}
