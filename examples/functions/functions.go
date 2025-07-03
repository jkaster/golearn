package functions

import "fmt"

// Functions demonstrates various types of functions in Go.
func Functions() {
	fmt.Println("\n--- Functions Example ---")

	// Simple function
	greet("Alice")

	// Function with return value
	sum := add(5, 3)
	fmt.Printf("Sum of 5 and 3 is: %d\n", sum)

	// Function with multiple return values
	quotient, remainder := divide(10, 3)
	fmt.Printf("10 divided by 3 is %d with a remainder of %d\n", quotient, remainder)

	// Variadic function
	total := sumAll(1, 2, 3, 4, 5)
	fmt.Printf("Sum of all numbers is: %d\n", total)

	// Anonymous function (closure)
	multiplier := 2
	multiply := func(x int) int {
		return x * multiplier
	}
	fmt.Printf("5 multiplied by %d is: %d\n", multiplier, multiply(5))

	// Function as a parameter
	process(10, 5, add)
	process(10, 5, subtract)
}

// greet is a simple function that takes a name and prints a greeting.
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// add is a function that takes two integers and returns their sum.
func add(a, b int) int {
	return a + b
}

// divide is a function that takes two integers and returns their quotient and remainder.
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// sumAll is a variadic function that takes a variable number of integers and returns their sum.
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// process is a function that takes two integers and a function as a parameter.
func process(a, b int, operation func(int, int) int) {
	result := operation(a, b)
	fmt.Printf("Operation result: %d\n", result)
}

// subtract is a helper function for process.
func subtract(a, b int) int {
	return a - b
}
