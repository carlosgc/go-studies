// Package importdecl shows how use import declarations in Go.
package main

// Importing multiple packages.
import (
	"fmt"                              // fmt is used in qualified identifiers to access exported identifiers.
	pkg1 "gostudies/importdecl/pkgone" // pkg1 is used in qualified identifiers to access exported identifiers.
	_ "gostudies/importdecl/pkgthree"  // Using blank identifier so package init function can be called.
	. "gostudies/importdecl/pkgtwo"    // All package's exported identifiers will be declared in this package.
)

func main() {

	// Calling function of fmt package.
	fmt.Println("fmt package imported.")

	// Calling function of pkgone package.
	pkg1.PkgOneGreeting()

	// Calling function of pkgtwo package.
	PkgTwoGreeting()
}
