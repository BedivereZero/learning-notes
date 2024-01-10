package main

import (
	"context"
	"log/slog"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.ContextWithFallback = true

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/int", func(c *gin.Context) {
		// var output chan int
		var output = make(chan int)
		go generateRandomIntSlow(c.Request.Context(), 8, output)

		var timer <-chan time.Time
		if d, err := time.ParseDuration(c.Query("timeout")); err == nil {
			slog.Info("set timeout", "duration", d)
			timer = time.After(d)
		}

		select {
		case <-timer:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "timeout"})
		case i := <-output:
			c.JSON(http.StatusOK, gin.H{"result": i})
		}

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func generateRandomIntSlow(ctx context.Context, maxRetry int, output chan<- int) error {
	t := time.NewTicker(time.Millisecond * 100)
	var retry int
	for {
		select {
		case <-ctx.Done():
			slog.Error("context is done", "error", ctx.Err())
			return ctx.Err()
		case <-t.C:
			if retry < maxRetry {
				retry++
				slog.Info("generating integer", "retry", retry)
				continue
			}
			output <- rand.Int()
			return nil
		}
	}
}
