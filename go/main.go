package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "%v\n", "hogehoge")
	})

	question := r.Group("question")
	{
		question.GET("", func(c *gin.Context) {
			c.String(http.StatusOK, "%v\n", c.Request.Method)
		})

		question.POST("", func(c *gin.Context) {
			c.String(http.StatusOK, "%v\n", c.Request.Method)
		})

		question.PUT("/", func(c *gin.Context) {
			c.String(http.StatusOK, "%v\n", c.Request.Method)
		})

		question.DELETE("/", func(c *gin.Context) {
			c.String(http.StatusOK, "%v\n", c.Request.Method)
		})
	}

	collection := r.Group("collection")
	{
		collection.GET("", func(c *gin.Context) {
			c.String(http.StatusOK, "%v\n", c.Request.Method)
		})

		collection.POST("", func(c *gin.Context) {
			c.String(http.StatusOK, "%v\n", c.Request.Method)
		})

		collection.PUT("/", func(c *gin.Context) {
			c.String(http.StatusOK, "%v\n", c.Request.Method)
		})

		collection.DELETE("/", func(c *gin.Context) {
			c.String(http.StatusOK, "%v\n", c.Request.Method)
		})
	}

	r.Run()
}
