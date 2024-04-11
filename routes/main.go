package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", nil)
	})

	api := server.Group("/api")
	api.POST("/products-scan", func(c *gin.Context) {
		// q := c.Request.URL.Query()
		// qrcode := q.Get("qrcode")
		// barcode := q.Get("barcode")
		b := c.Request.Body
		if qrcode == "" && barcode == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "qrcode or barcode is required",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
	api.PUT("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
}
