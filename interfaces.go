package main

import "fmt"

// Speaker is an interface that defines a method for speaking.
type Speaker interface {
	Speak() string
}

// Dog struct implements the Speaker interface.
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("Woof! My name is %s", d.Name)
}

// Cat struct also implements the Speaker interface.
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("Meow! My name is %s", c.Name)
}

// SaySomething takes a Speaker interface as an argument.
func SaySomething(s Speaker) {
	fmt.Println(s.Speak())
}

// interfaces demonstrates how to define and use interfaces in Go, including polymorphism and the empty interface.
func interfaces() {
	// Create instances of Dog and Cat.
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	// Call SaySomething with both Dog and Cat, demonstrating polymorphism.
	SaySomething(dog)
	SaySomething(cat)

	// An empty interface can hold values of any type.
	var i interface{}
	i = 42
	fmt.Printf("Empty interface holding an int: %v (type: %T)\n", i, i)
	i = "hello"
	fmt.Printf("Empty interface holding a string: %v (type: %T)\n", i, i)
}
