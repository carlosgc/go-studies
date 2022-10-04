package main

import "fmt"

/*
	Type Alias Declaration

	Syntax:
		type typeName = sourceType

	With type alias declaration the typeName and sourceType denote the same type.
*/

// Type aliases can be declared at package level

// Single type alias declaration

type type1 = int

// Multiple type aliases declaration

type (
	type2 = int
	type3 = string
	type4 = type1
)

/*
	Type Definition

	Syntax:
		type typeName sourceType

	With type definition the new defined type and its respective source type in type definitions are two distinct types.
*/

// Type definition can be declared at package level

// Single type definition

type type5 int

// Multiple type definition

type (
	type6 int
	type7 string
	type8 type1
)

var (
	t1 type1
	t2 type2
	t3 type3
	t4 type4
	t5 type5
	t6 type6
	t7 type7
	t8 type8
)

func typeDeclarationExemples() {

	fmt.Println("\n=========================\nType Declaration Examples\n=========================\n")

	// Type aliases and type definition can be declared at function level.

	type fType1 = int // Type alias
	type fType2 int   // Type definition

	var (
		ft1 fType1
		ft2 fType2
	)

	fmt.Println("The type of aliases")
	fmt.Printf("type1: \t%T\n", t1) // type1:  int
	fmt.Printf("type2: \t%T\n", t2) // type2:  int
	fmt.Printf("type3: \t%T\n", t3) // type3:  string
	fmt.Printf("type4: \t%T\n", t4) // type4:  int
	fmt.Printf("fType1: %T\n", ft1) // fType1: int

	fmt.Println("\nThe type of new defined types")
	fmt.Printf("type5: \t%T\n", t5) // type5:  main.type5
	fmt.Printf("type6: \t%T\n", t6) // type6:  main.type6
	fmt.Printf("type7: \t%T\n", t7) // type7:  main.type7
	fmt.Printf("type8: \t%T\n", t8) // type8:  main.type8
	fmt.Printf("fType2: %T\n", ft2) // fType2: main.fType2
}
