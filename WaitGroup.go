package main

import (
	"fmt"
	"sync"
)

func calculateLines(c int, wg *sync.WaitGroup) {
	fmt.Println("Called")
	if c != 0 {
		wg.Add(1)
		calculateLines(c - 1, wg)
	}
	wg.Done()
}

func main() {
	var wg = &sync.WaitGroup{}

	wg.Add(1)
	go calculateLines(3, wg)
	fmt.Println("Waiting")
	wg.Wait()
}
