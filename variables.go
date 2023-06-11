package main

import "fmt"

func variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2 // You can declare multiple variables at once.
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // Variables declared without a corresponding initialization are zero-valued.
	fmt.Println(e)

	f := "apple" // The := syntax is shorthand for declaring and initializing a variable.
	fmt.Println(f)
}
