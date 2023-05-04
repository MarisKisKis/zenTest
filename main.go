package main

import (
	"zenTest/cmd/api"
	"zenTest/cmd/database"
)

func init() {
	database.NewPostgreSQLClient()
}
func main() {
	r := api.SetupRouter()
	r.Run("localhost:8080")
}
