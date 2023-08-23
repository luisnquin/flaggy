package main

import "github.com/luisnquin/flaggy"

func main() {
	// Declare variables and their defaults
	positionalValue := "defaultValue"

	// Add the positional value to the parser at position 1
	flaggy.AddPositionalValue(&positionalValue, "test", 1, true, "a test positional value")

	// Parse the positional value
	flaggy.Parse()

	// Use the value
	print(positionalValue)
}
