package main

/*
	Named constants, or constants, are all boolean, numeric or string values.
	Constantes can be untyped or typed.

	The type of an untyped constant is the same as the value bound to it.

	Untyped constant declaration sytanx:
		const constantName = value

	Typed constant declaration sytanx:
		const constantName type = value

	Unnamed constants are the literals bound to named constants. Ex.: 1, "A0"
*/

/*
	Constants can be declared in at package level.
	In this case they are called package-level or global contants.

	The declaration orders of two package-level constants are not important.
	In the following example, the declaration orders of PKG_C4 and PKG_C9 can be exchanged.
*/

// Single constant declarations.

const PKG_C1 = 0           // untyped constant
const PKG_C2 int = 0       // typed constant
const PKG_C3 = string("0") // typed constant

// Multiple constant declarations.

// Multiple inline constant declarations.

const PKG_C4, PKG_C5, PKG_C6 = 0, PKG_C9, int(0) // PKG_C4 and PKG_C5 are untyped constants and PKG_C6 is a typed constant
const PKG_C7, PKG_C8 int = 0, 0                  // typed constants

// Multiple grouped constant declarations.

const (
	PKG_C9         = true
	PKG_C10 string = "0"
	PKG_C11        = string("0")
)

/*
	In a group-style constant declaration, the first constant specification needs to be bound to a value.
	Incomplete constant specifications (with no value bound) will be bound with the same value as the
	first preceding complete constant specification.

	In group-style constant declarations, each line containing a `=` symbol is known as constant specification.
*/

const (
	PKG_C12 = 1
	PKG_C13 // PKG_C13 = 1

	PKG_C14 = "A"
	PKG_C15 // PKG_C15 = "A"
	PKG_C16 // PKG_C16 = "A"

	PKG_C17, PKG_C18 = 1, 2
	PKG_C19, _       // PKG_C19, _ = 1, 2
)

/*
	Constant declaration can use the iota constant generator feature of Go.
	iota is a predeclared constant which can only be used in other constant declarations.

	The value of iota will be reset to 0 at the first constant specification of each constant declaration.
	iota value will increase 1 in each constant specification.
*/

const PKG_C20 = iota // iota = 0, PKG_C20 = 0

const (
	PKG_C21          = iota       // iota = 0, PKG_C21 = iota = 0
	PKG_C22                       // iota = 1, PKG_C22 = iota = 1
	_                             // iota = 2
	PKG_C23          = iota + 1   // iota = 3, PKG_C23 = iota + 1 = 4
	PKG_C24                       // iota = 4, PKG_C24 = iota + 1 = 5
	PKG_C25, PKG_C26 = iota, iota // iota = 5, PKG_C25 = PKG_C26 = iota = 5
)

const (
	PKG_C27 = iota // iota = 0, PKG_C27 = iota = 0
	PKG_C28        // iota = 1, PKG_C28 = iota = 1
)

func constantExamples() {

	/*
		Constants can be declared in at function level.
		In this case they are called function-level or local contants.

		The declaration orders of two function-level constants are important.
		In the following example, the declaration orders of FNC_C1 and FNC_C3 can't be exchanged.
	*/

	const FNC_C1 = 0
	// const FNC_C2 = FNC_C3 	// error!!
	const FNC_C2 = 0
	const FNC_C3 = FNC_C1
}
