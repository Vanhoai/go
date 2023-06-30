package main

import "fmt"

func rangeGoLang() {

    numbers := []int{2, 3, 4}
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range numbers {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }
}