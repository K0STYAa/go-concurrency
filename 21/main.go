// Реализовать паттерн «адаптер» на любом примере.
package main

import "fmt"

// Target provides an interface with which the system should work.
type Target interface {
	Request() string
}

// Adaptee implements system to be adapted.
type Adaptee struct {
}

// NewAdapter is the Adapter constructor.
func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

// SpecificRequest implementation.
func (a *Adaptee) SpecificRequest() string {
	return "Request"
}

// Adapter implements Target interface and is an adapter.
type Adapter struct {
	*Adaptee
}

// Request is an adaptive method.
func (a *Adapter) Request() string {
	return a.SpecificRequest()
}

func main() {
	adapter := NewAdapter(&Adaptee{})

	req := adapter.Request()

	if req != "Request" {
		fmt.Println("Expected req = ", req)
	}

	fmt.Println(req)
}
