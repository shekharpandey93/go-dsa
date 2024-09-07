package main

import (
	"fmt"
	"sync"
)

func main() {
	//c := make(chan int)
	//var wg sync.WaitGroup
	//wg.Add(2)
	//go odd(c, &wg)
	//go even(c, &wg)
	//wg.Wait()

	var wg sync.WaitGroup
	hiChan := make(chan bool, 1)
	helloChan := make(chan bool, 1)

	wg.Add(2)

	go printHi(hiChan, helloChan, &wg)
	go printHello(hiChan, helloChan, &wg)

	hiChan <- true // Start the sequence by signaling the "hi" goroutine

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

func printHi(hiChan, helloChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		<-hiChan          // Wait for the signal from the "hello" goroutine
		fmt.Println("hi") // Print "hi"
		helloChan <- true // Signal the "hello" goroutine to print "hello"
	}
}

func printHello(hiChan, helloChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		<-helloChan          // Wait for the signal from the "hi" goroutine
		fmt.Println("hello") // Print "hello"
		hiChan <- true       // Signal the "hi" goroutine to print "hi"
	}
}
