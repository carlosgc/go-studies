package pkgtwo

import "fmt"

// The init() function is invoked before the main() function.
func init() {
	fmt.Println("pkgtwo imported!")
}

// Prints a greetings message for this file.
func PkgTwoGreeting() {
	fmt.Println("Hello pkgtwo.go")
}
