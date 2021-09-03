package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type Middleware struct {
	app  *firebase.App
	auth *auth.Client
}

func NewMiddleware(firebaseCredentialsFilePath string) *Middleware {
	opt := option.WithCredentialsFile(firebaseCredentialsFilePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Panic("firebase init err: ", err)
	}
	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Panic("firebase init err: ", err)
	}
	return &Middleware{app: app, auth: auth}
}

func (m *Middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// jwtを取得
		authorization := c.Request.Header.Get("Authorization")
		idToken := strings.Replace(authorization, "Bearer ", "", 1)

		token, err := m.auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			// JWT が無効なら失敗判定
			c.String(http.StatusUnauthorized, "error verifying ID token: %v\n", err)
			c.Abort()
		} else {
			log.Printf("Verified ID token: %v\n", token)
			c.Next()
		}
	}
}
