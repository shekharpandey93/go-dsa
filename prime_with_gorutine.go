package main

import (
	"fmt"
	"sync"
)

// Function to check if a number is prime
func isPrimeVal(num int) bool {
	if num == 1 {
		return false
	}
	if num == 2 || num == 3 {
		fmt.Println(num)
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	for idx := 5; idx*idx <= num; idx = idx + 6 {
		if num%idx == 0 || num%(idx+2) == 0 {
			return false
		}
		fmt.Println(idx)
	}
	return true
}

// Function to generate prime numbers up to a limit and send them through a channel
func generatePrimes(limit int, primeChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := 2; num <= limit; num++ {
		if isPrimeVal(num) {
			primeChan <- num
		}
	}
	close(primeChan) // Close the channel after all primes are sent
}
func printPrimeVal(primeChan chan int) {
	for prime := range primeChan {
		fmt.Println(prime)
	}
}

func main() {
	primeLimit := 50
	primeChan := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)

	// Start a goroutine to generate primes
	go generatePrimes(primeLimit, primeChan, &wg)

	// Print the prime numbers received from the channel
	go printPrimeVal(primeChan)

	wg.Wait() // Wait for the prime generation goroutine to finish
}
