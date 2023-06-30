package main

import "fmt"

type rectangle struct {
	width, height int
}

func (r *rectangle) area() int {
	return r.width * r.height
}

func (r rectangle) perim() int {
	return 2*r.width + 2*r.height
}

type user struct {
	name string
	age  int
}

func (u *user) sayHello() string {
	return "Say Hello"
}

func method() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	u := user{name: "Jon", age: 20}
	fmt.Println(u.sayHello())
}
