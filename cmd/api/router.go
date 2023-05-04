package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", home)
	r.POST("/postgres/users", postUser)
	r.POST("/sign/hmacsha512", postSha512)
	r.GET("/postgres/users/:id", getUser)
	r.DELETE("/postgres/users/:id", deleteUser)
	return r
}
