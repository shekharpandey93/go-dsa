package main

import (
	"fmt"
	"sync"
)

func printTableFive(num int, group *sync.WaitGroup) {
	defer group.Done()
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d X %d = %d\n", num, i, num*i)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go printTableFive(5, &wg)
	wg.Wait()
}
