//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

var set = wire.NewSet(
	NewEvent,
	wire.Bind(new(Greeter), new(*GreeterHello)),
	wire.Struct(new(GreeterHello), "Message"),
	NewMessage,
)

func InitializeEvent(phrase string) (Event, error) {
	wire.Build(set)
	return Event{}, nil
}

func InitializePrinter(filename string) (*Printer, func(), error) {
	wire.Build(NewPrinter, provideFile)
	return nil, nil, nil
}
