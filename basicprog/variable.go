package main

import "fmt"

func variableExamples() {

	fmt.Println("\n=================\nVariable Examples\n=================\n")

	// Single variable declaration.
	// The variable is initialized with the zero type value (default value).

	var v1 int // v1 will be initialized with 0 (int type zero value).
	fmt.Println("Variable v1:", v1)

	// Declaring a variable with an initial value

	var v2 int = 20 // v2 will be initialized with 20.
	fmt.Println("Variable v2:", v2)

	// When declaring a variable with no type specification an with an initial value
	// Go will automatically infer the variable type

	var v3 = 30 // v3 will be an int type variable initialized with 30
	fmt.Println("Variable v3:", v3)

	// Multiple variables declaration

	var v4, v5 int // v4 e V5 will be initialized with 0 (int type zero value).
	fmt.Println("Variable v4:", v4)
	fmt.Println("Variable v5:", v5)

	var v6, v7 int = 60, 70 // v6 will be initialized with 60 and v7 will be initialized with 70.
	fmt.Println("Variable v6:", v6)
	fmt.Println("Variable v7:", v7)

	// Multiples variables declaration of different types

	var (
		v8 int    = 80   // v8 will be initialized with 80
		v9 string = "90" // v9 will be initialized with "90"
	)
	fmt.Println("Variable v8:", v8)
	fmt.Println("Variable v9:", v9)

	// Multiples variables declaration with type inference

	var v10, v11 = 100, "110" // v10 will be an int type variable initialized with 100 and v11 will be string type variable initialized with "110".
	fmt.Println("Variable v10:", v10)
	fmt.Println("Variable v11:", v11)

	// Short hand declaration

	v12 := 120 // v12 will be an int type variable initialized with 120
	fmt.Println("Variable v12:", v12)

	// Multiple short hand declaration
	// Short hand syntax can only be used when at least one of the variables on the left side of := is newly declared

	v13, v14 := 130, "140" // v13 will be an int type variable initialized with 130 and v14 will be string type variable initialized with "140".
	fmt.Println("Variable v13:", v13)
	fmt.Println("Variable v14:", v14)
}
