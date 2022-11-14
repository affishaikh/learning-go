package main

import "fmt"
import (
	"learning-go/sample/another"
	"learning-go/sample/another1"
)

func main() {
	//fmt.Println(sample.Random())
	fmt.Println(another.AnotherRandom())
	fmt.Println(another1.Random())
}
