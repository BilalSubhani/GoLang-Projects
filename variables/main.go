package main

import "fmt"

// Package-level (global) variable declarations
var globalInt int = 100
var globalFloat float64 = 3.1415
var globalBool = true   // type inferred
var globalString string // zero value (empty string)

// Multiple variable declaration
var (
	a    int
	b, c float32
	d, e = 1, "hello"
)

func main() {
	// Function-level variable declarations and initializations
	var localInt int = 42
	var localFloat float64 = 2.718
	var localBool bool // zero value (false)
	var localString string = "GoLang"

	// Short variable declaration (:=)
	x := 10
	y := 3.14
	isActive := false
	message := "Short declaration!"

	// Multiple short variable declaration
	p, q, r := 1, 2.5, "multi"

	// Constants (not variables, but important for data types)
	const pi = 3.14159
	const greeting string = "Hello, constant!"

	// Print all variables
	fmt.Println("Package-level variables:", globalInt, globalFloat, globalBool, globalString)
	fmt.Println("Multiple package-level:", a, b, c, d, e)
	fmt.Println("Function-level variables:", localInt, localFloat, localBool, localString)
	fmt.Println("Short declarations:", x, y, isActive, message)
	fmt.Println("Multiple short declarations:", p, q, r)
	fmt.Println("Constants:", pi, greeting)

	// Type conversion
	var intToFloat float64 = float64(localInt)
	fmt.Println("Type conversion (int to float64):", intToFloat)

	// Zero values demonstration
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string
	fmt.Println("Zero values:", zeroInt, zeroFloat, zeroBool, zeroString)
}
