package main

import "fmt"

// Person is a struct that represents a person with a Name and Age.
type Person struct {
	Name string
	Age  int
}

// Greet method for the Person struct.
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}

// Kilometer is a custom type based on float64.
type Kilometer float64

// Miles method for the Kilometer type.
func (k Kilometer) Miles() float64 {
	return float64(k) * 0.621371
}

func customTypes() {
	// Create an instance of the Person struct.
	p := Person{Name: "Alice", Age: 30}
	fmt.Printf("Person: %+v\n", p)
	fmt.Println(p.Greet())

	// Create another instance, demonstrating field access.
	p2 := Person{"Bob", 25} // Field names can be omitted if order is maintained
	fmt.Printf("Person 2 Name: %s, Age: %d\n", p2.Name, p2.Age)

	// Anonymous struct
	anon := struct {
		Color string
		Size  int
	}{"Red", 10}
	fmt.Printf("Anonymous struct: %+v\n", anon)

	// Use the custom Kilometer type.
	var distance Kilometer = 10.0
	fmt.Printf("Distance in Kilometers: %v km\n", distance)
	fmt.Printf("Distance in Miles: %v miles\n", distance.Miles())

	// Type alias (not a new type, just an alternative name)
	type MyString = string
	var myStr MyString = "This is a custom string alias."
	fmt.Printf("MyString: %s (type: %T)\n", myStr, myStr)
}
