package main

import (
	"arrayLib/myarray"
	"fmt"
)

func main() {
	fmt.Println("--- Array Declaration ---")
	myarray.DeclareArray()

	fmt.Println("--- Array Initialization ---")
	myarray.InitializeArray()

	fmt.Println("--- Accessing and Modifying Array ---")
	myarray.AccessArray()

	fmt.Println("--- Iterating Array ---")
	myarray.IterateArray()

	fmt.Println("--- Passing Array as Parameter ---")
	arr := [3]int{7, 8, 9}
	myarray.ArrayAsParameter(arr)

	fmt.Println("--- Returning Array from Function ---")
	returned := myarray.ReturnArray()
	fmt.Println("Returned array:", returned)

	fmt.Println("--- Multi-dimensional Array ---")
	myarray.MultiDimensionalArray()
}
