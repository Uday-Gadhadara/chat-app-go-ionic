package main

import (
	"chat/backend/database"
	"chat/backend/router"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	app := gin.Default()
	r := app.Group("/api")
	{
		router.IndexRouter(r)
	}
	app.Run()
}
