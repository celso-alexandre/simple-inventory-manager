package main

import (
	"fmt"

	"github.com/celso-alexandre/simple-inventory-manager/db"
	"github.com/celso-alexandre/simple-inventory-manager/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	_, err := db.Connect()

	if err != nil || db.DB == nil {
		fmt.Println(err)
		fmt.Println("Error connecting to the database")
		return
	}
	fmt.Println("Connected to the database")

	server := gin.Default()
	server.LoadHTMLGlob("templates/*.html")

	routes.RegisterRoutes(server)

	server.Run(":3333")
}
