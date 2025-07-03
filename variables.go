package main

import "fmt"

func variables() {
	// Declare a variable of type string with the name 'a' and assign it the value "initial"
	var a = "initial"
	fmt.Println(a)

	// Declare two integer variables, b and c, and assign them the values 1 and 2 respectively
	var b, c int = 1, 2
	fmt.Println(b, c)

	// The compiler will infer the type of an initialized variable
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued.
	// For example, the zero value for an int is 0.
	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable,
	// e.g. for var f string = "apple" in this case.
	f := "apple"
	fmt.Println(f)
}
