package main

import (
	"postly/configs"
	"postly/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	configs.Postgresql()
	router := gin.Default()
	routes.ApiRoutes(router)
	router.Run(":8080")
}
