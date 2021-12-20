package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().UTC()
	ca, err := time.LoadLocation("Canada/Atlantic")
	if err != nil {
		panic(err)
	}
	now = now.In(ca)
	fmt.Println(now.Format(time.RFC3339Nano))
}
