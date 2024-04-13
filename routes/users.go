package routes

import (
	"fmt"
	"net/http"

	"github.com/celso-alexandre/simple-inventory-manager/middlewares"
	"github.com/celso-alexandre/simple-inventory-manager/models"
	"github.com/celso-alexandre/simple-inventory-manager/utils"
	"github.com/gin-gonic/gin"
)

func userLogin(c *gin.Context) {
	var u models.User
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	foundUser, err := models.FindUserByUsername(u.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	u.Id = foundUser.Id

	if foundUser.Id <= 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found"})
		return
	}

	if foundUser.Password == "" && u.Password != "" {
		err := u.UpdatePassword()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		foundUser, err = models.FindUserByUsername(u.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if foundUser.Password == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password must be provided"})
		return
	}

	err = utils.ComparePassword(foundUser.Password, u.Password)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	jwtToken := utils.GenerateJwtToken(utils.JwtPayload{Username: u.Username, UserId: u.Id})
	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}

func userSignup(c *gin.Context) {
	var u models.User
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jwtPayload := middlewares.RetrieveAuthPayload(c)
	u.UpdatedByUserId = jwtPayload.UserId
	err = u.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, u)
}

// func getUsers(c *gin.Context) {
// 	users, err := models.FindAllUsers()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, users)
// }

// func getUserById(c *gin.Context) {
// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	user, err := models.FindUserById(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }

// func updateUser(c *gin.Context) {
// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	var u models.User
// 	err = c.BindJSON(&u)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	u.Id = id
// 	err = u.Update()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, u)
// }

// func deleteUser(c *gin.Context) {
// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	err = models.DeleteUser(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.Status(http.StatusNoContent)
// }
