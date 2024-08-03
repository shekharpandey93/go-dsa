package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go odd(c, &wg)
	go even(c, &wg)
	wg.Wait()
}

func even(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		<-c
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func odd(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		c <- 1
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}
