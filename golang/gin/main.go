package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"

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
	r := gin.New()
	r.Use(
		gin.Logger(),
		gin.CustomRecovery(RecoveryToResponse()),
	)

	r.POST("/panic-if-zero", PanicIfZero)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// 尝试向 response 写入 recovery 信息
func RecoveryToResponse() gin.RecoveryFunc {
	return func(c *gin.Context, err any) {
		var buf = make([]byte, 1024)
		runtime.Stack(buf, false)

		c.JSON(http.StatusInternalServerError, gin.H{
			"err":   err,
			"stack": buf,
		})
	}
}

// PanicIfZero 当接收到 0 时 panic
func PanicIfZero(c *gin.Context) {
	var req = struct {
		Value  int       `json:"value,omitempty"`
		Reader io.Reader `json:"reader,omitempty"`
	}{}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	io.ReadAll(req.Reader)

	if req.Value == 0 {
		panic("req.value == 0")
	}

	c.JSON(http.StatusOK, &req)
}
