package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/celso-alexandre/simple-inventory-manager/models"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", nil)
	})

	api := server.Group("/api")
	api.POST("/products-scan", func(c *gin.Context) {
		var product models.Product
		err := c.BindJSON(&product)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}
		if product.Barcode == "" && product.Id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "id or barcode is required",
			})
			return
		}
		err = product.SaveScan()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			return
		}
		c.JSON(http.StatusOK, product)
	})

	api.PUT("/products/:id", func(c *gin.Context) {
		var product models.Product
		err := c.BindJSON(&product)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid body.id",
			})
			return
		}
		product.Id = id
		err = product.Update()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			return
		}
		c.JSON(http.StatusOK, product)
	})
}
