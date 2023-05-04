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

// saving user to database and returning its id
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
