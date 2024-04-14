package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/celso-alexandre/simple-inventory-manager/server/middlewares"
	"github.com/celso-alexandre/simple-inventory-manager/server/models"
	"github.com/gin-gonic/gin"
)

func productScan(c *gin.Context) {
	var product models.Product
	err := c.BindJSON(&product)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	if product.Barcode == "" && product.Uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "uuid or barcode is required",
		})
		return
	}
	jwtPayload := middlewares.RetrieveAuthPayload(c)
	product.UpdatedByUserId = jwtPayload.User.UserId
	err = product.SaveScan()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, product)
}

func productUpdate(c *gin.Context) {
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
	jwtPayload := middlewares.RetrieveAuthPayload(c)
	product.UpdatedByUserId = jwtPayload.User.UserId
	err = product.Update()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, product)
}
