package main

import (
	"fmt"
	"sync"
	"time"
)

func anotherBomb(wg *sync.WaitGroup) {
	for i := 1; i < 11; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Boom!!")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go anotherBomb(&wg)
	go anotherBomb(&wg)
	go anotherBomb(&wg)
	wg.Wait()
}
