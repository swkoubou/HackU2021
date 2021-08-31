package main

import (
	"net/http"

	"example.com/handle"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "%v\n", "hogehoge")
	})

	question := r.Group("/question")
	{
		question.GET("/all", handle.GetAllQuestionsHandler)

		question.GET("/:questionID", handle.GetQuestionsHandler)

		question.POST("", func(c *gin.Context) {
			c.String(http.StatusOK, "%v\n", c.Request.Method)
		})

		question.PUT("/:questionID", func(c *gin.Context) {
			c.String(http.StatusOK, "%v\n", c.Request.Method)
		})

		question.DELETE("/:questionID", func(c *gin.Context) {
			c.String(http.StatusOK, "%v\n", c.Request.Method)
		})
	}

	collection := r.Group("/collection")
	{
		collection.GET("/all", handle.GetAllCollectionsHandler)

		collection.GET("/:collectionID", handle.GetCollectionHandler)

		collection.POST("", func(c *gin.Context) {
			c.String(http.StatusOK, "%v\n", c.Request.Method)
		})

		collection.PUT("", func(c *gin.Context) {
			c.String(http.StatusOK, "%v\n", c.Request.Method)
		})

		collection.DELETE("", func(c *gin.Context) {
			c.String(http.StatusOK, "%v\n", c.Request.Method)
		})
	}

	r.Run()
}
