package main

import (
	"arrays/arraylib"
	"fmt"
)

func main() {
	fmt.Println("--- Array Declaration ---")
	myarray.DeclareArray()

	fmt.Println("\n\n--- Array Initialization ---")
	myarray.InitializeArray()

	fmt.Println("\n\n--- Accessing and Modifying Array ---")
	myarray.AccessArray()

	fmt.Println("\n\n--- Iterating Array ---")
	myarray.IterateArray()

	fmt.Println("\n\n--- Passing Array as Parameter ---")
	arr := [3]int{7, 8, 9}
	myarray.ArrayAsParameter(arr)

	fmt.Println("\n\n--- Returning Array from Function ---")
	returned := myarray.ReturnArray()
	fmt.Println("Returned array:", returned)

	fmt.Println("\n\n--- Multi-dimensional Array ---")
	myarray.MultiDimensionalArray()
}
