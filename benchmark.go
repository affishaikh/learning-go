package main

import (
	"fmt"
	"strconv"
)

func makeRange(min, max int) []string {
	a := make([]string, max-min+1)
	for i := range a {
		a[i] = strconv.Itoa(min + i)
	}
	return a
}

func main() {
	arr := makeRange(0, 100000)
	var s, sep string
	//for i := 1; i < len(os.Args); i++ {
	for _, j := range arr {
		s += sep + j
		sep = " "
	}

	fmt.Println(s)
}
//
//func main() {
//	arr := makeRange(0, 100000)
//	fmt.Println(strings.Join(arr[0:], " "))
//}