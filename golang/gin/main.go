package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewLogMiddleware(name string, n int) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("%s: start, n=%d", name, n)
		defer log.Printf("%s: end, n=%d", name, n)

		if n > 1 {
			return
		}

		c.Request = c.Request.WithContext(NewContextWithOperator(c.Request.Context(), &Operator{Name: "Remilia"}))

		c.Next()
	}
}

func process(ctx context.Context) {
	log.Println(OperatorFromContext(ctx))
}

func Echo(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		body = []byte(fmt.Sprintf("error: %v", err))
	}

	process(c)

	c.JSON(http.StatusOK, gin.H{
		"body":          string(body),
		"header":        c.Request.Header,
		"remoteAddress": c.Request.RemoteAddr,
		"method":        c.Request.Method,
		"path":          c.Request.RequestURI,
	})
}

func registerAPI(r gin.IRouter) gin.IRouter {
	// a := r.Group("a", NewLogMiddleware("a", 0))
	// b := a.Group("b", NewLogMiddleware("b", 1))
	// c := b.Group("c", NewLogMiddleware("c", 2))

	r.Use(WithOperator())

	r.GET("a/b/c", Echo)

	return r
}

func main() {
	r := gin.Default()

	a := r.Group("a", NewLogMiddleware("a", 0))
	b := a.Group("b", NewLogMiddleware("b", 1))
	c := b.Group("c", NewLogMiddleware("c", 2))
	c.Any("echo", Echo)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
