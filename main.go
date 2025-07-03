package main

import (
	"fmt"

	"golang/examples/binarytree"
	"golang/examples/collections"
	"golang/examples/customtypes"
	"golang/examples/datatypes"
	"golang/examples/goroutines"
	"golang/examples/interfaces"
	"golang/examples/scope"
	"golang/examples/variables"
)

// main is the entry point of the Go learning project.
// It calls various example functions to demonstrate Go concepts.
func main() {
	fmt.Println("Hello, World!")

	fmt.Println("\n--- Variables Example ---")
	variables.Variables()

	fmt.Println("\n--- Scope Example ---")
	scope.Scope()

	fmt.Println("\n--- Data Types Example ---")
	datatypes.DataTypes()

	fmt.Println("\n--- Custom Types Example ---")
	customtypes.CustomTypes()

	fmt.Println("\n--- Interfaces Example ---")
	interfaces.Interfaces()

	fmt.Println("\n--- Binary Tree Example ---")
	binarytree.BinaryTreeExample()

	fmt.Println("\n--- Goroutines Example ---")
	goroutines.Goroutines()

	fmt.Println("\n--- Collections Example ---")
	collections.Collections()
}
