package datatypes

import "fmt"

// DataTypes demonstrates various built-in data types in Go, including booleans, integers, floats, and strings.
// It also shows zero values and type inference.
func DataTypes() {
	// Boolean type
	var b bool = true
	fmt.Printf("Boolean: %v (type: %T)\n", b, b)

	// Numeric types (integers)
	var i int = 10
	var i8 int8 = 127 // Max value for int8
	var u16 uint16 = 65535 // Max value for uint16
	fmt.Printf("Integer: %v (type: %T)\n", i, i)
	fmt.Printf("Int8: %v (type: %T)\n", i8, i8)
	fmt.Printf("Uint16: %v (type: %T)\n", u16, u16)

	// Numeric types (floating-point)
	var f32 float32 = 3.14
	var f64 float64 = 3.1415926535
	fmt.Printf("Float32: %v (type: %T)\n", f32, f32)
	fmt.Printf("Float64: %v (type: %T)\n", f64, f64)

	// String type
	var s string = "Hello, Go!"
	fmt.Printf("String: %v (type: %T)\n", s, s)

	// Zero values
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string
	fmt.Printf("Zero values: int=%v, float=%v, bool=%v, string=\"%v\"\n", zeroInt, zeroFloat, zeroBool, zeroString)

	// Type inference with := (short declaration)
	x := 42
	y := 3.14159
	z := "Go is fun"
	fmt.Printf("Inferred types: x=%v (%T), y=%v (%T), z=%v (%T)\n", x, x, y, y, z, z)
}