package main

import "fmt"

// The init() function is invoked before the main() function.
func init() {

	fmt.Println("init(): The init() function in srcone.go file.")
}

// Prints a greetings message for this file.
func hellosrcone() {

	fmt.Println("Hello srcone.go!")
}
