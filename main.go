package main

import (
	"github.com/celso-alexandre/simple-inventory-manager/db"
	"github.com/celso-alexandre/simple-inventory-manager/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	db.Connect()

	server := gin.Default()
	server.LoadHTMLGlob("templates/*.html")

	routes.RegisterRoutes(server)

	server.Run(":3333")
}
