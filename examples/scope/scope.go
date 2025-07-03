package scope

import "fmt"

// globalVar is a global variable, accessible to any function in the package.
var globalVar = "I'm a global variable"

// Scope demonstrates variable scoping rules in Go: global, local, and block-level.
func Scope() {
	// localVar is a local variable, accessible only within the scope function.
	var localVar = "I'm a local variable"

	fmt.Println(globalVar)
	fmt.Println(localVar)

	if true {
		// blockVar is a block-level variable, accessible only within this if block.
		var blockVar = "I'm a block-level variable"
		fmt.Println(blockVar)
	}

	// The following line would cause a compile error because blockVar is not accessible here.
	// fmt.Println(blockVar)
}
