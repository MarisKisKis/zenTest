package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"zenTest/cmd/model"
)

func postIncr(c *gin.Context) {
	var incr model.Incr
	if err := c.ShouldBindJSON(&incr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	client := rClient()
	ctx := context.Background()
	err := client.Set(ctx, incr.Key, incr.Value, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, incr.Key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(incr.Key, val)
	return
}

func rClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,
	})

	return client
}
