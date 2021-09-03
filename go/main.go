package main

import (
	"log"
	"net/http"

	"example.com/handle"
	"example.com/manage"
	"example.com/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	firebaseCredentialsFilePath := "firebase-adminsdk.json"
	middleware := middleware.NewMiddleware(firebaseCredentialsFilePath)

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

	auth := r.Group("/private", middleware.Auth())
	{
		auth.GET("/mydata", func(c *gin.Context) {
			c.String(http.StatusOK, "%v\n", "isLogin")
		})
	}

	r.GET("/check", func(c *gin.Context) {
		db, err := manage.NewDBConnection()
		if err != nil {
			log.Println("NewDBConnection(): ", err)
			return
		}

		if err := db.Ping(); err != nil {
			log.Println("db.Ping(): ", err)
		} else {
			log.Println("db.Ping(): success")
		}
	})

	r.Run()
}
