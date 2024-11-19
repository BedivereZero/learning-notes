// main.go
package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	alert := js.Global().Get("alert")
	alert.Invoke("Hello World!")

	fmt.Println("Hello, world!")
}
