package main

import (
	"fmt"
	"strconv"
	"sync"
)

func worker1(wg *sync.WaitGroup, cs chan string, i int) {
	defer wg.Done()
	cs <- "worker" + strconv.Itoa(i)
	if i == 9 {
		close(cs)
	}
}

func monitorWorker1(cs chan string) {
	for i := range cs {
		fmt.Println(i)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	cs := make(chan string)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker1(wg, cs, i)
	}

	monitorWorker1(cs)
	wg.Wait()
}
