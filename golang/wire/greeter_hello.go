package main

type GreeterHello struct {
	Message Message
}

var _ Greeter = &GreeterHello{}

func NewGreeterHello(m Message) *GreeterHello {
	return &GreeterHello{
		Message: m,
	}
}

// Greet implements Greeter.
func (g *GreeterHello) Greet() Message {
	return g.Message
}

// Grumpy implements Greeter.
func (*GreeterHello) Grumpy() bool {
	return false
}
