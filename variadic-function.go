package main

import "fmt"

// the same spread operator in JavaScript
func sum(numbers ...int) {
    fmt.Print(numbers, " ")
    total := 0

    for _, num := range numbers {
        total += num
    }
    fmt.Println(total)
}

func variadicFunction() {

    sum(1, 2)
    sum(1, 2, 3)

    numbers := []int{1, 2, 3, 4}
	// the same rest operator in JavaScript
    sum(numbers...)
}