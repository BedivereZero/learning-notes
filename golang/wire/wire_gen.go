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
	greeterHello := &GreeterHello{
		Message: message,
	}
	event, err := NewEvent(greeterHello)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

// wire.go:

var set = wire.NewSet(
	NewEvent, wire.Bind(new(Greeter), new(*GreeterHello)), wire.Struct(new(GreeterHello), "Message"), NewMessage,
)
