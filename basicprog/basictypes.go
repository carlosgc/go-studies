package main

import "fmt"

var b1 bool // Boolean type variable initialized as false

var s1 string // String type variable initialized as empty string ("")

// Implementation-specific integer numeric types variables
var (
	i1    int
	ui1   uint
	uiptr uintptr
)

// Integer numeric types variables of different sizes and initialized as zero value (0)
var (
	i8  int8  // signed 8 bit size integer type
	ui8 uint8 // unsigned 8 bit size integer type

	i16  int16  // signed 16 bit size integer type
	ui16 uint16 // unsigned 16 bit size integer type

	i32  int32  // signed 32 bit size integer type
	ui32 uint32 // unsigned 32 bit size integer type

	i64  int64  // signed 64 bit size integer type
	ui64 uint64 // unsigned 64 bit size integer type
)

// Floating-point numeric types types variables of different sizes and initialized as zero value (0.0)
var (
	f32 float32 // 32 bit size floating-point type
	f64 float64 // 64 bit size floating-point type
)

// Complex numeric types types variables of different sizes and initialized as zero value (0.0)
var (
	c64  complex64  // 64 bit size complex type
	c128 complex128 // 128 bit size complex type
)

var by1 byte // Byte type is an alias for uint8 type

var r1 rune // Rune type is an alias for uint8 int32

func basicTypesExamples() {

	fmt.Println("\n=========================\nBasic Types Examples\n=========================\n")

	//The forms to write integer values
	fmt.Println("Integer literals")
	fmt.Println("Decimal form (base 10): 14 =", 14)
	fmt.Println("Binary form (base 2): 0b1110 =", 0b1110)
	fmt.Println("Octal form (base 8): 016 =", 016)
	fmt.Println("Hexadecimal form (base 16): 0xe =", 0xe)

	//The forms to write float point values
	fmt.Println("\nFloat point literals")
	fmt.Println("Decimal form")
	fmt.Println("0.23 =", 0.23)
	fmt.Println(".23 =", 0.23)
	fmt.Printf("1. = %.2f\n", 1.)
	fmt.Printf("1.23e2 = %.2f\n", 1.23e2)
	fmt.Printf("1.23e-2 = %.4f\n", 1.23e-2)

	fmt.Println("\nHexadecimal form")
	fmt.Printf("0x.8p2 = (1 + 8/16)*2² = %.2f\n", 0x.8p2)
	fmt.Printf("0x.8p-2 = (1 + 8/16)/2² = %.3f\n", 0x.8p-2)
	fmt.Printf("0x8p2 = 8.0*2² = %.2f\n", 0x8p2)
	fmt.Printf("0x8p-2 = 8.0/2² = %.2f\n", 0x8p-2)
	fmt.Printf("0x8.4p2 = (1 + 4/16)*2² = %.2f\n", 0x8.4p2)
	fmt.Printf("0x8.4p-2 = (1 + 4/16)/2² = %.4f\n", 0x8.4p-2)

	//The forms to write complex values
	fmt.Println("\nComplex literals")
	fmt.Println("1 + 2i =", 1+2i)
	fmt.Println("1.5 + .2i =", 1.5+.2i)
	fmt.Println("2i - 1 =", 2i-1)
	fmt.Println("1.2i =", 1.2i)

	fmt.Println("\nDecimal form (base 10): 5i =", 5i)
	fmt.Println("Binary form (base 2): 0b101i =", 0b101i)
	fmt.Println("Octal form (base 8): 0o5i =", 0o5i)
	fmt.Println("Hexadecimal form (base 16): 0x5i =", 0x5i)

	// Improving values readability
	fmt.Println("\nReadabilty improveness")
	fmt.Println("1_2 =", 1_2)
	fmt.Println("11_22 =", 11_22)
	fmt.Println("111_222 =", 111_222)

	fmt.Println("\nDecimal form (base 10): 2_333_444 =", 2_333_444)
	fmt.Println("Binary form (base 2): 0b_10_00_01 =", 0b_10_00_01)
	fmt.Println("Octal form (base 8): 0o1_2_3 =", 0o1_2_3)
	fmt.Println("Hexadecimal form (base 16): 0xA_B_C =", 0xA_B_C)

	//The forms to write rune values
	fmt.Println("\nRune literals")
	fmt.Println("Character form: 'a' =", 'a')
	fmt.Println("Octal form: '\\141' =", '\141')
	fmt.Println("Two hex digits Hexadecimal form: '\\x61' =", '\x61')
	fmt.Println("Four hex digits Hexadecimal form: '\\u0061' =", '\u0061')
	fmt.Println("Eight hex digits Hexadecimal form: '\\U00000061' =", '\U00000061')

	//The forms to write string values
	fmt.Println("\nString literals")
	fmt.Println("\"Hello world!\" =", "Hello world!")
	fmt.Println("'Hello world!' =", `Hello world!`)
	fmt.Println("\"Hello\\nworld!\" =", "Hello\nworld!")
	fmt.Println("'Hello\nworld!' =", `Hello
world!`)
}
