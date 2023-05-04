package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", home)
	r.POST("/redis/incr", postIncr)
	r.POST("/sign/hmacsha512", postSha512)
	r.POST("/postgres/users", postUser)
	r.GET("/postgres/users/:id", getUser)
	r.DELETE("/postgres/users/:id", deleteUser)
	return r
}
