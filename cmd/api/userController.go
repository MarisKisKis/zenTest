package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zenTest/cmd/database"
)

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Go home",
	})
	return
}

func postUser(c *gin.Context) {
	var user database.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	res, err := database.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": res,
	})
	return
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	user, err := database.GetUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
	return
}
