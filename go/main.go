package main

import (
	"context"
	"log"
	"net/http"

	"example.com/handle"
	"example.com/manage"
	"example.com/middleware"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func main() {

	// firebase init
	firebaseCredentialsFilePath := "firebase-adminsdk.json"
	opt := option.WithCredentialsFile(firebaseCredentialsFilePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Panic("firebase init err: ", err)
	}
	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Panic("firebase init err: ", err)
	}

	middleware := middleware.NewMiddleware(app, auth)
	handle := handle.NewHandle(app, auth)

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

	private := r.Group("/private", middleware.Auth())
	{
		private.GET("/mydata", func(c *gin.Context) {
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
