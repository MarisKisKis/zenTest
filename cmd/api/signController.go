package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"zenTest/cmd/model"
)

func postSha512(c *gin.Context) {
	var sign model.Sign
	if err := c.ShouldBindJSON(&sign); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	fmt.Println(sign)
	return
}
