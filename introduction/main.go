/*
   package main

   'package main' marks this file as part of the main package.
   This declaration is required for building an *executable* Go program.

   The main package is special because:
   - It defines the entry point of the application.
   - It must include a 'main()' function.
   - The program starts executing from the 'main()' function.

   Without 'package main', the Go compiler will not produce an executable binary.

   --- Package Types in Go ---

   1. Executable Package:
      - package main
      - Used for standalone programs intended to be executed directly.
      - Must contain a 'main()' function.
      - Compiles into a runnable binary.

   2. Reusable Package:
      - package mylib (or any custom name)
      - Used to encapsulate reusable code.
      - Can be imported into other packages.
      - Cannot act as an entry point and doesn’t include a 'main()' function.

	--- Importing Packages in Go ---

	In Go, packages are imported to use their functionality.

   import "fmt": This imports the "fmt" package from the Go standard library.
   The "fmt" package (short for 'format') provides formatted I/O functions
   and is commonly used for printing to the console.

   Key functions include:
   - fmt.Println(...)  → Prints with a newline.
   - fmt.Printf(...)   → Supports formatted strings (like printf in C).
   - fmt.Sprint(...)   → Returns a formatted string without printing it.

   Example:
     fmt.Println("Hello, Go!")

   "fmt" is essential for debugging, logging, and general output formatting.
*/

// Package declaration
package main

// Importing the fmt package
import "fmt"

// Declare the main function: main is the entry point of the application
func main() {
	fmt.Println("Hello, World!")
}
