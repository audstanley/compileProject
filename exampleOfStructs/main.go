package main

import (
	"fmt"
)

// fuck this
type operation struct {
	s    string
	x, y int
}

// this interface is available to the operations
// struct
type T interface {
	runOperation()
}

// The run function is implecitely called from
// an operation struct, runs the operation
// and returns a struct.
func (o operation) runOperation() int {
	if o.s == "+" {
		return o.x + o.y
	} else if o.s == "-" {
		return o.x - o.y
	} else if o.s == "*" {
		return o.x * o.y
	} else if o.s == "/" {
		if o.y == 0 {
			panic("You tried to divide by zero.")
		}
		return o.x / o.y
	}
	return 0
}

func main() {

	var o = operation{"*", 3, 4}
	fmt.Printf("%d", o.runOperation())

}
