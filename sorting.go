package main

import (
	"fmt"
	"sort"
)

type FileMetadata struct {
	Name string
	NumberOfLines int
}

func main() {
	a := []FileMetadata{{"", 80}, {"", 1}, {"", 4}, {"", 9}, {"", 7}, {"", 36}}
	b := make([]FileMetadata, len(a))
	copy(b, a)
	fmt.Println("b--------->", b)
	fmt.Println("a--------->", a)
	sort.Slice(b, func(i, j int) bool {
		return b[i].NumberOfLines > b[j].NumberOfLines
	})

	fmt.Println(b)
}
