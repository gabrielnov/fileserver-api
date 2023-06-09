package controllers

import "github.com/gin-gonic/gin"

func Upload(c *gin.Context) {
	c.String(200, "upload working")
}

func Download(c *gin.Context) {
	c.String(200, "download working")
}
