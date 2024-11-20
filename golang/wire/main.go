package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type Message string

func NewMessage(phrase string) Message {
	return Message(phrase)
}

type Greeter struct {
	Message Message
	Grumpy  bool
}

func NewGreeter(m Message) Greeter {
	var grumpy bool = time.Now().Second()%2 == 0
	return Greeter{Message: m, Grumpy: grumpy}
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

type Event struct {
	Greeter Greeter
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	// message := NewMessage()
	// greeter := NewGreeter(message)
	// event := NewEvent(greeter)
	// event.Start()

	e, err := InitializeEvent("Hello, world!")
	if err != nil {
		log.Fatalf("create event fail: %v", err)
	}
	e.Start()
}
