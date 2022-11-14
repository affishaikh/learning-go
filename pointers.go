package main

import "fmt"

func main() {
	a := 42
	p := &a
	fmt.Println(a)
	fmt.Println(*p)
	*p = 82
	fmt.Println(*p)
	fmt.Println(a)
}
