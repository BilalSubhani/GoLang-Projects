package myarray

import "fmt"

// DeclareArray demonstrates array declaration and default values
func DeclareArray() {
	var arr [5]int
	fmt.Println("Declared array:", arr)
}

// InitializeArray demonstrates array initialization
func InitializeArray() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Initialized array:", arr)
}

// AccessArray demonstrates accessing and modifying array elements
func AccessArray() {
	arr := [3]string{"Go", "is", "fun"}
	fmt.Println("Original array:", arr)
	arr[1] = "can be"
	fmt.Println("Modified array:", arr)
}

// IterateArray demonstrates iterating over an array
func IterateArray() {
	arr := [4]int{10, 20, 30, 40}
	fmt.Print("Iterating array:")
	for i, v := range arr {
		fmt.Printf(" Index %d: %d;", i, v)
	}
	fmt.Println()
}

// ArrayAsParameter demonstrates passing an array to a function
func ArrayAsParameter(arr [3]int) {
	fmt.Println("Array received as parameter:", arr)
}

// ReturnArray demonstrates returning an array from a function
func ReturnArray() [2]float64 {
	arr := [2]float64{3.14, 2.71}
	return arr
}

// MultiDimensionalArray demonstrates 2D arrays
func MultiDimensionalArray() {
	var matrix [2][3]int
	matrix[0][1] = 7
	matrix[1][2] = 9
	fmt.Println("2D array (matrix):", matrix)
}
