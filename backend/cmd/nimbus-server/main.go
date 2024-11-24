package main

import (
	"github.com/anirudhi89/nimbus/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		routes.InitializeRoutes(api)
	}

	r.Run(":8080")
}
