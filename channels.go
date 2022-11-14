package main

import "fmt"

func putANumber(c chan int) {
	for i := 0; i <= 20; i++ {
		c <- i
	}
}

func main() {
	c := make(chan int)
	go putANumber(c)
	for i := 0; i <= 20; i++ {
		v := <- c
		fmt.Println(v)
	}
}
