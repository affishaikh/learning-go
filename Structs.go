package main

import "fmt"

type Person struct {
	name string
	age int
}

func change(p Person) {
	fmt.Println(p)
	p.age = 100
	fmt.Println(p)
}

func main() {
	p := Person{"r", 10}
	change(p)
	fmt.Println(p)
}
