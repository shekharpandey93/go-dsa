package main

import (
	"fmt"
	"sync"
)

func printNumber(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d\n", i)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go printNumber(&wg)
	wg.Wait()
}
