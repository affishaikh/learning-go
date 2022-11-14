package main

import "fmt"

func printSomething() {
	fmt.Println("Work")
}

func main() {
	defer fmt.Print("End of the program!")
	go printSomething()
}
