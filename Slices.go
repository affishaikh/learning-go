package main

import "fmt"

func lowBoundHighBound() {
	s := []int{10, 8, 62, 48, 82}
	printSlice(s)

	s = s[:3]
	printSlice(s)
	s = s[:4]
	printSlice(s)
	s = s[1:4]
	printSlice(s)
	s = s[1:4]
	printSlice(s)
	s = s[1:8]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func rangeExp() {
	a := [][]int {{1,2}, {3,4}}

	for _, nArr := range a {
		nArr[1] = 10
	}
	fmt.Println(a)
}

func appendingToSlice() {
	arr := [6]int {10, 8, 42, 30, 68, 1}
	s := arr[:4]
	printSlice(s)
	newArr := append(s, 11, 1)
	printSlice(s[1:6])
	printSlice(newArr)
	printSlice(arr[:])
}

func main() {
	//lowBoundHighBound()
	//appendingToSlice()
	rangeExp()
}
