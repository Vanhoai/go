package main

import "fmt"

func getValues() (int, int) {
    return 3, 7
}

func multipleValueReturn() {

    a, b := getValues()
    fmt.Println(a)
    fmt.Println(b)

    _, c := getValues()
    fmt.Println(c)
}