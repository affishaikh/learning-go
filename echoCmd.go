package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	//for i := 1; i < len(os.Args); i++ {
	for _, j := range os.Args[1:] {
		s += sep + j
		sep = " "
	}

	fmt.Println(s)
}

//func main() {
//	fmt.Println(strings.Join(os.Args[1:], " "))
//}