package api

import (
	"context"
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
	// connection to Redis
	client := rClient()
	ctx := context.Background()
	err := client.Set(ctx, incr.Key, incr.Value, 0).Err()
	if err != nil {
		panic(err)
	}
	// incrementing by key
	incrByKey, err := client.Incr(ctx, incr.Key).Result()
	if err != nil {
		panic(err)
	}
	// returning value in json format
	var value incrValue
	value.Value = incrByKey
	c.JSON(http.StatusOK, value)
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

type incrValue struct {
	Value int64 `json:"value"`
}
