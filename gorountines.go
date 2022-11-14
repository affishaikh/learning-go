package main

import (
	"fmt"
	"sync"
	"time"
)

func bomb() {
	for i := 1; i < 11; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Boom!!")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		bomb()
		wg.Done()
	}()
	go func() {
		bomb()
		wg.Done()
	}()
	go func() {
		bomb()
		wg.Done()
	}()
	wg.Wait()
}
