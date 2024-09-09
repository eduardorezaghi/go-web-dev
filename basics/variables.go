package main

import "fmt"

// Outer scope
const (
	GLOBAL_FIRST = iota
	GLOBAL_SECOND
	GLOBAL_THIRD
)

func main() {
	// Inner scope
	var intVar int = 10
	var floatVar float64 = 10.5
	var stringVar string = "Hello, World!"

	fmt.Printf("intVar: %d\nfloatVar: %f\nstringVar: %s\n", intVar, floatVar, stringVar)
}
