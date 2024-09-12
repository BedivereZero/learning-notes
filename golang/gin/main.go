package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/BedivereZero/learning-notes/golang/gin/object"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)

	r := gin.Default()

	object.RegisterAPI(r)
	registerPing(r)

	if err := r.Run(); err != nil {
		log.Printf("run server fail: %v", err)
	}
}
