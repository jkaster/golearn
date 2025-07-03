package main

import (
	"fmt"

	"golearn/examples/binarytree"
	"golearn/examples/collections"
	"golearn/examples/customtypes"
	"golearn/examples/datatypes"
	"golearn/examples/goroutines"
	"golearn/examples/interfaces"
	"golearn/examples/scope"
	"golearn/examples/variables"
	"golearn/examples/functions"
)

// main is the entry point of the Go learning project.
// It calls various example functions to demonstrate Go concepts.
func main() {
	fmt.Println("Hello, World!")

	variables.Variables()

	scope.Scope()

	datatypes.DataTypes()

	customtypes.CustomTypes()

	interfaces.Interfaces()

	binarytree.BinaryTreeExample()

	goroutines.Goroutines()

	collections.Collections()

	functions.Functions()
}
