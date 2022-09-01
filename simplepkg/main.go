// A Go program has to have a main package to use as a entry point.

// Package simpleprog shows how to create a simple program with Go.
package main

import "fmt"

// The init() function is invoked before the main() function.
func init() {

	fmt.Println("init(): A function to init the package.")
}

// The main() function is a special function that is the entry point of an executable program.
func main() {

	fmt.Println("Hello Simple Package.")

	hellosrcone()

	hellosrctwo()
}

// Each package source file can declare multiple functions named as init(). At run time, each init() function will be sequentially invoked once (before invoking the main() entry function).
func init() {

	fmt.Println("Second init(): Cool!! Packages can have more than one init.")
}
