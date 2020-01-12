package main

import (
	"net/http"
	"os"
	. "github.com/gorefa/gin-jwt/handler"

	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gorefa/log"
)


func main() {
	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "8081"
	}

	err := JWTinit()
	if err != nil {
		log.Fatalf(err,"jwt init error")
	}

	r.POST("/login", AuthMiddleware.LoginHandler)

	r.NoRoute(AuthMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Infof("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", AuthMiddleware.RefreshHandler)
	auth.Use(AuthMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", HelloHandler)
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("http start error",err)
	}
}