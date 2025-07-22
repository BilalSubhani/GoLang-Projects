package main

import (
	"fmt"
)

func main() {
	// 1. Declaration and Initialization
	var m1 map[string]int // nil map, cannot assign
	fmt.Println("m1 (nil map):", m1)

	// m1["a"] = 1 // This would panic (assignment to nil map)

	// Using make
	m2 := make(map[string]int)
	m2["one"] = 1
	m2["two"] = 2
	fmt.Println("m2 (make):", m2)

	// Literal initialization
	m3 := map[string]int{"a": 10, "b": 20}
	fmt.Println("m3 (literal):", m3)

	// 2. Accessing values
	val := m3["a"]
	fmt.Println("m3[\"a\"]:", val)

	// 3. Checking key existence
	v, ok := m3["b"]
	fmt.Println("Key 'b' exists?", ok, "Value:", v)
	v2, ok2 := m3["c"]
	fmt.Println("Key 'c' exists?", ok2, "Value:", v2)

	// 4. Iterating over map
	fmt.Println("Iterating m2:")
	for k, v := range m2 {
		fmt.Printf("%s => %d\n", k, v)
	}

	// 5. Deleting a key
	delete(m2, "one")
	fmt.Println("m2 after delete:", m2)

	// 6. Map length
	fmt.Println("len(m3):", len(m3))

	// 7. Map of slices
	mapOfSlices := map[string][]int{
		"evens": {2, 4, 6},
		"odds":  {1, 3, 5},
	}
	fmt.Println("mapOfSlices:", mapOfSlices)

	// 8. Slice of maps
	sliceOfMaps := []map[string]int{
		{"x": 1},
		{"y": 2},
	}
	fmt.Println("sliceOfMaps:", sliceOfMaps)

	// 9. Map with struct values
	type person struct {
		age int
	}
	people := map[string]person{
		"Alice": {30},
		"Bob":   {25},
	}
	fmt.Println("Map with struct values:", people)

	// 10. Map with struct keys (not allowed if struct contains slice/map/func)

	type key struct {
		id int
	}
	m4 := make(map[key]string)
	m4[key{1}] = "one"
	fmt.Println("Map with struct key:", m4)

	// 11. Passing map to function (reference type)
	modifyMap := func(m map[string]int) {
		m["changed"] = 999
	}
	modifyMap(m3)
	fmt.Println("m3 after modifyMap:", m3)

	// 12. Maps are reference types
	m5 := m3
	m5["a"] = 111
	fmt.Println("m3 after m5[\"a\"]=111:", m3)

	// 13. Initializing nil map with make
	m1 = make(map[string]int)
	m1["foo"] = 42
	fmt.Println("m1 after make:", m1)

	// 14. Nested maps
	nested := map[string]map[string]int{
		"outer": {"inner": 1},
	}
	fmt.Println("Nested map:", nested)

	// 15. Deleting non-existent key (no panic)
	delete(m2, "notfound")
	fmt.Println("m2 after deleting non-existent key:", m2)

	// 16. Map with interface{} values
	m6 := map[string]interface{}{
		"num": 1,
		"str": "hello",
		"arr": []int{1, 2, 3},
	}
	fmt.Println("Map with interface{} values:", m6)

	// 17. Comparing maps (not allowed except nil check)
	// m2 == m3 // invalid operation
	fmt.Println("m2 == nil?", m2 == nil)

	// 18. Map capacity (not exposed, but can hint with make)
	m7 := make(map[int]int, 100)
	fmt.Println("m7 (with capacity hint):", m7)

	// 19. Concurrent map access (unsafe, not shown here)
	// Use sync.Map for safe concurrent access
}
