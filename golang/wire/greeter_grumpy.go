package main

type GreeterGrumpy struct {
	Message Message
}

var _ Greeter = &GreeterGrumpy{}

// Greet implements Greeter.
func (g *GreeterGrumpy) Greet() Message {
	return g.Message
}

// Grumpy implements Greeter.
func (*GreeterGrumpy) Grumpy() bool {
	return true
}
