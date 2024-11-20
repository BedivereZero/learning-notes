//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

var set = wire.NewSet(
	NewEvent,
	NewGreeterHello,
	NewMessage,
)

func InitializeEvent(phrase string) (Event, error) {
	wire.Build(set)
	return Event{}, nil
}
