package main

import (
	"errors"
	"fmt"
	"log"
)

type Message string

func NewMessage(phrase string) Message {
	return Message(phrase)
}

type Greeter interface {
	Greet() Message
	Grumpy() bool
}

type Event struct {
	Greeter Greeter
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy() {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	e, err := InitializeEvent("Hello, world!")
	if err != nil {
		log.Fatalf("create event fail: %v", err)
	}
	e.Start()
}
