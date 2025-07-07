package main

import (
	"fmt"
	"functions/mylib"
)

func main() {
	// Swap two strings
	a, b := "hello", "world"
	swappedA, swappedB := mylib.SwapStrings(a, b)
	fmt.Println("Swapped:", swappedA, swappedB)

	// Add two numbers
	sum := mylib.AddNumbers(3, 5)
	fmt.Println("Sum:", sum)

	// Check if a number is even
	n := 4
	fmt.Printf("%d is even? %v\n", n, mylib.IsEven(n))

	// Check number sign
	fmt.Println("Sign of -7:", mylib.NumberSign(-7))
	fmt.Println("Sign of 0:", mylib.NumberSign(0))
	fmt.Println("Sign of 9:", mylib.NumberSign(9))

	// Day type using switch
	fmt.Println("Monday is a:", mylib.DayType("Monday"))
	fmt.Println("Sunday is a:", mylib.DayType("Sunday"))
	fmt.Println("Funday is a:", mylib.DayType("Funday"))

	// Sum up to n using for loop
	fmt.Println("Sum up to 5:", mylib.SumUpTo(5))
}
