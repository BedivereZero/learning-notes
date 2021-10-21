package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	direcotry := os.Getenv("DIRECTORY")
	log.Println("directory:", direcotry)
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(direcotry))))
}
