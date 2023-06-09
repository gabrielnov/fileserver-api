package webserver

import (
	"github.com/gin-gonic/gin"
	"uploads-api/internal/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	v1 := r.Group("/api/v1")

	v1.POST("/upload", controllers.Upload)
	v1.GET("/download", controllers.Download)

	return r
}
