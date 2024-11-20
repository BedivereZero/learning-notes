package main

type GreeterHello struct {
	Message Message
}

var _ Greeter = &GreeterHello{}

// Greet implements Greeter.
func (g *GreeterHello) Greet() Message {
	return g.Message
}

// Grumpy implements Greeter.
func (*GreeterHello) Grumpy() bool {
	return false
}
