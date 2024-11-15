package main

import (
	"log"
	"os"
)

func main() {
	log.Default().SetFlags(log.Ltime | log.Lmicroseconds | log.Lshortfile)

	if err := NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
