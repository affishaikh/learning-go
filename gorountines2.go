package main

import (
	"fmt"
	"sync"
)

func numbers(name string, wg *sync.WaitGroup) {
	for i := 1; i < 100; i++ {
		fmt.Printf("%v %v \n", name, i)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go numbers("first", &wg)
	go numbers("second", &wg)
	wg.Wait()
}
