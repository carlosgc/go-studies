package pkgone

import "fmt"

// The init() function is invoked before the main() function.
func init() {
	fmt.Println("pkgone imported!")
}

// Prints a greetings message for this file.
func PkgOneGreeting() {
	fmt.Println("Hello pkgone.go")
}
