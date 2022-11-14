package main

import "fmt"

func main() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%v %v %v %v %v %v", &a[0], &a[1], &a[2], &a[3], &a[4], &a[5])
}
