package main

import (
	"fmt"
)

func main() {
	// 1. Declaration and Initialization
	// Declare a slice with zero value (nil)
	var s1 []int
	fmt.Println("s1 (nil slice):", s1, "len:", len(s1), "cap:", cap(s1))

	// Declare and initialize with values
	s2 := []int{1, 2, 3}
	fmt.Println("s2 (initialized):", s2)

	// 2. Creating a slice with make
	s3 := make([]int, 5) // length 5, capacity 5
	fmt.Println("s3 (make):", s3, "len:", len(s3), "cap:", cap(s3))

	s4 := make([]int, 2, 10) // length 2, capacity 10
	fmt.Println("s4 (make with cap):", s4, "len:", len(s4), "cap:", cap(s4))

	// 3. Slicing an array
	arr := [5]int{10, 20, 30, 40, 50}
	s5 := arr[1:4] // elements 20, 30, 40
	fmt.Println("s5 (slice of array):", s5)

	// 4. Slicing a slice
	s6 := s2[1:]
	fmt.Println("s6 (slice of slice):", s6)

	// 5. Appending elements
	s7 := append(s2, 4, 5)
	fmt.Println("s7 (after append):", s7)

	// Appending another slice
	s8 := append(s2, s5...)
	fmt.Println("s8 (append another slice):", s8)

	// 6. Copying slices
	s9 := make([]int, len(s2))
	copy(s9, s2)
	fmt.Println("s9 (copy of s2):", s9)

	// 7. Iterating over slices
	fmt.Print("Iterating s7:")
	for i, v := range s7 {
		fmt.Printf(" [%d]=%d", i, v)
	}
	fmt.Println()

	// 8. Modifying elements
	s7[0] = 100
	fmt.Println("s7 after modification:", s7)

	// 9. Length and capacity
	fmt.Println("len(s7):", len(s7), "cap(s7):", cap(s7))

	// 10. Nil vs empty slice
	var s10 []int  // nil
	s11 := []int{} // empty
	fmt.Println("s10 (nil):", s10, s10 == nil)
	fmt.Println("s11 (empty):", s11, s11 == nil)

	// 11. Multi-dimensional slices
	matrix := [][]int{
		{1, 2},
		{3, 4, 5},
	}
	fmt.Println("matrix (2D slice):", matrix)

	// 12. Removing an element (e.g., index 2)
	s12 := []int{1, 2, 3, 4, 5}
	idx := 2
	s12 = append(s12[:idx], s12[idx+1:]...)
	fmt.Println("s12 after removing index 2:", s12)

	// 13. Passing slices to functions
	printSlice := func(s []int) {
		fmt.Println("Inside function, slice:", s)
	}
	printSlice(s12)

	// 14. Slices as reference types
	s13 := []int{10, 20, 30}
	s14 := s13
	s14[0] = 99
	fmt.Println("s13 (after s14[0]=99):", s13) // s13 is affected

	// 15. Growing slice capacity
	s15 := []int{}
	for i := 0; i < 10; i++ {
		s15 = append(s15, i)
		fmt.Printf("After append %d: len=%d cap=%d\n", i, len(s15), cap(s15))
	}

	// 16. Slice of structs
	type person struct {
		name string
		age  int
	}
	people := []person{
		{"Alice", 30},
		{"Bob", 25},
	}
	fmt.Println("Slice of structs:", people)

	// 17. Slices with interface{}
	var anySlice []interface{} = []interface{}{1, "two", 3.0}
	fmt.Println("Slice of interface{}:", anySlice)

	// 18. Comparing slices (deep equality)
	// Use reflect.DeepEqual for content comparison
	// (not shown here to avoid importing reflect)
	// 19. Slices in Go are not comparable with == except for nil check

	// 20. Slicing beyond capacity (will panic)
	// Uncommenting below will panic:
	// fmt.Println(s2[0:10])

	// 21. Appending to nil slice is valid
	var s16 []int
	s16 = append(s16, 1)
	fmt.Println("Appending to nil slice:", s16)
}
