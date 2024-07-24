package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name": "World",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/submit", func(c *gin.Context) {

		username := c.PostForm("username")
		email := c.PostForm("email")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"email":    email,
		})
	})
}
