package main

import "fmt"

type person struct {
	name string
	age  int
}

type linux struct {
	lll int
}

func main() {
	p := person{name: "hhhhh"}
	fmt.Printf("han的类型是 %v", p)
	g := 111.11
	fmt.Printf("%T", g)
	
}
