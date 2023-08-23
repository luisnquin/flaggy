package main

import "github.com/luisnquin/flaggy"

func main() {
	// Declare variables and their defaults
	stringFlag := "defaultValue"

	// Add a flag
	flaggy.String(&stringFlag, "f", "flag", "A test string flag")

	// Parse the flag
	flaggy.Parse()

	// Use the flag
	print(stringFlag)
}
