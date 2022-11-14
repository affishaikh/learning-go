package main

import "fmt"

type Point struct {
	x int
	y int
}

type Board struct {
	upperBound Point
	lowerBound Point
}

func main() {
	board := Board{Point{0, 0}, Point{8, 8}}

	fmt.Println(board)
}
