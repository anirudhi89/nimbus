package router

// Routes for handling different API requests

import (
	"github.com/anirudhi89/nimbus/internal/user"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.RouterGroup) {
	version1 := r.Group("/v1")
	{
		exampleInitializeRoutes(version1)
	}
}

func exampleInitializeRoutes(r *gin.RouterGroup) {
	exampleRoutes := r.Group("/example")
	{
		exampleRoutes.GET("/:id", user.GetIDTestCall) // localhost:8080/api/v1/example/:id
	}
}
