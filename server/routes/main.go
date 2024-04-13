package routes

import (
	"net/http"

	"github.com/celso-alexandre/simple-inventory-manager/server/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", nil)
	})

	api := server.Group("/api")

	api.POST("/login", userLogin)

	api.Use(middlewares.AuthMiddleware())
	api.POST("/signup", userSignup)
	api.POST("/products-scan", productScan)
	api.PUT("/products/:id", productUpdate)
}
