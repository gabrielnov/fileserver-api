package webserver

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.New()

	v1 := r.Group("/api/v1")

	v1.POST("/upload", nil)
	v1.GET("/download", nil)
	return r
}
