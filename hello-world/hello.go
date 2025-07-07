package main

import "fmt"

func main() {
	name := "Bilal"
	fmt.Println("Hello,", name)

	firstName := "Bilal"
	lastName := "Subhani"
	age := 26
	design := "Software Engineer"

	// var firstName, lastName, design string = "Bilal", "Subhani", "Software Engineer"

	fmt.Printf("Hello, I am %s %s. I am %d years old, and currently I am working as a %s.\n", firstName, lastName, age, design)
}
