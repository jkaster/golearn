package main

import (
	"fmt"
	"math"
)

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

	// Constants are declared using the const keyword.
	const s string = "constant"
	fmt.Println(s)

	// A const statement can appear anywhere a var statement can.
	const n = 500000000

	// Constant expressions perform arithmetic with arbitrary precision.
	const d2 = 3e20 / n
	fmt.Println(d2)

	// A numeric constant has no type until it’s given one, such as by an explicit cast.
	fmt.Println(int64(d2))

	// A number can be given a type by using it in a context that requires one,
	// for example, a variable assignment or function call.
	// For example, math.Sin expects a float64.
	fmt.Println(math.Sin(n))
}
