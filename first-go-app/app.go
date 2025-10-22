// Package declaration. Every Go file starts with a package declaration.
// The 'main' package is a special package in Go that indicates that this file is part of an executable program.
// When you run the program, the Go runtime looks for the main package to start execution.
package main

// You must have at least one package per file.
// You can have multiple files in a package.
// You can have multiple packages in a single project.
// The idea of a package is to group related code together.

// Importing the fmt package to use its functions for formatted I/O operations.
// It is built-in package in Go standard library.
import "fmt"

// The main function is the entry point of the Go program.
// When you run the program, execution starts from the main function.
// There can only be one main function in a package.
func main() {
	//We are calling the Println function from the fmt package to print "Hello, World!" to the console.
	fmt.Println("Hello, World!")
	// it accepts backticks for multi-line strings
	fmt.Println(`This is a multi-line string.`)
}

//Modules are a way to group related Go packages together.
// They allow you to manage dependencies and versioning for your Go projects.

// go mod init <module-name>
// This command initializes a new module in your project with the specified module name.
// It creates a go.mod file that tracks your module's dependencies.

// go mod tidy
// This command cleans up your go.mod file by adding any missing module requirements
// and removing any that are no longer necessary. It ensures that your module's dependencies
// are in sync with the code in your project.

//go build
// This command compiles the Go source code files in your module and produces an executable binary.
// The binary will have the same name as the module or the directory containing the main package.
// To execute the binary, you can run it from the command line. For example, on Unix-like systems, you can use:
// ./<binary-name>
// On Windows, you would use:
// <binary-name>.exe

//go run <file-name>.go
// This command compiles and runs the specified Go source code file in a single step.
// It is useful for quickly testing and executing small Go programs without creating a binary file.
