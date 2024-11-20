// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeEvent(phrase string) (Event, error) {
	message := NewMessage(phrase)
	greeter := NewGreeterHello(message)
	event, err := NewEvent(greeter)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

// wire.go:

var set = wire.NewSet(
	NewEvent,
	NewGreeterHello,
	NewMessage,
)
