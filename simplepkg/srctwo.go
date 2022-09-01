package main

import "fmt"

// The init() function is invoked before the main() function.
func init() {

	fmt.Println("init(): The init() function in srctwo.go file.")
}

// Prints a greetings message for this file.
func hellosrctwo() {

	fmt.Println("Hello srctwo.go!")
}
