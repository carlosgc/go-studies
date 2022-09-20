# Study notes

## Packages

* Go programs are organized into packages. They are constructed by linking together packages.

* A package is a collection of source files in the same directory that are compiled together. An implementation require that all source files for a package inhabit the same directory.

* A folder corresponds to a code package, and vice versa. The folder containing the source files of a code package is called the folder of the package.

* The source files in a folder (not including subfolders) must belong to the same package. It is impossible to have more than one package in one directory.

* A package is constructed from one or more source files that together declare constants, types, variables and functions belonging to the package and which are accessible in all files of the same package.

* Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.

* Each source file consists of a package clause defining the package to which it belongs.

```Go
package PackageName
```

* A set of files sharing the same PackageName form the implementation of a package.

* A sub-package is a package located inside a parent package. It's implemented as a package directory in the parent package directory

* Go has two types of packages:
    
    * Program packages (or command packages) - packages named with `main` and containing the `main` entry functions. Each Go program should contain one and only one program package.
    
    * Library packages - packages other than program packages.

* By convention, the package name and the package folder name are the same for library packages. For program packages they are different.

* Each package source file can declare multiple functions named as `init`. At run time, each `init` function will be sequentially invoked once (before invoking the `main` entry function).

* All package-level variables declared in a package are initialized before any `init` function declared in the same package is invoked.

## Import Declarations

* Import declaration is the way to use code and functionalities of other packages in your own program.

* The import names an identifier (PackageName) to be used for access and an ImportPath that specifies the package to be imported.

* The effects of import declaration are the follows:
    
    * States that the source file containing the declaration depends on functionality of the imported package.

    * Enables access to exported identifiers of that package.

* The way exported identifiers are accessed varies depending on how an import declaration is written.

* When PackageName is used in qualified identifiers to access exported identifiers of the package within the importing source file.

```Go
    import "PackageName"
    PackageName.FunctionName()
```

* When PackageName is omitted, it defaults to the identifier specified in the package clause of the imported package.

```Go
    import m "PackageName"
    m.FunctionName()
```

* When an explicit period (.) is used, all the package's exported identifiers will be declared in the importing source file's and must be accessed without a qualifier

```Go
    import . "PackageName"
    FunctionName()
```

* A package can't import itself, directly or indirectly.

* It's not possible import a package and without referring to any of its exported identifiers.

* It's possible to use a blank identifier as explicit package name to import it just for its side-effects (initialization)

```Go
    import _ "PackageName"
```

* Multiples packages can be imported using a import declaration with many packages:

```Go
    import (
        "PackageOne"
        "PackageTwo"
        "PackageThree"
    )
```

## Blank imports

* The importname in the full form import declaration can be the blank identifier (_). 

* Importing source files can't use the exported code elements in anonymously imported packages.

* The purpose of anonymous imports is to initialize the imported packages (each of init functions in the anonymously imported packages will be called once).

## Import paths

* An import path is a string used to import a package. A package's import path is its module path joined with its subdirectory within the module.

* An import path can describe how to obtain the package source code using a revision control system such as Git or Mercurial. The go tool uses this property to automatically fetch packages from remote repositories.

* A package whose import path containing an `internal` folder name is viewed as a special package. It can only be imported by the packages in and under the direct parent directory of the `internal` folder. For example, package `.../a/b/c/internal/d/e/f` and `.../a/b/c/internal` can only be imported by the packages whose import paths have a `.../a/b/c` prefix.

* Go’s convention is that the package name is the same as the last element of the import path.