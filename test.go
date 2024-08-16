package main

import (
	"fmt"
	"sync"
)

func main() {
	// todo chanel
	// wg 2
	c := make(chan int)
	printChant := make(chan int, 100)
	var wg sync.WaitGroup
	wg.Add(3)
	go odd(c, printChant, &wg)
	go even(c, printChant, &wg)
	go printMyVal(printChant, &wg)
	wg.Wait()
}

func odd(c, printChant chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		<-c
		if i%2 == 0 {
			printChant <- i
			//fmt.Println("odd:",i)
		}
	}
}

func even(c, printChant chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		c <- 1
		if i%2 == 1 {
			printChant <- i
			//fmt.Println("even",i)
		}
	}
}

func printMyVal(c chan int, wg *sync.WaitGroup) {
	//defer wg.Done()
	for {
		data := <-c
		fmt.Println("data", <-c)
		if data == 100 {
			wg.Done()
		}
	}
}

//func removeArray() {
//
//	mapDat := make(map[int]string)
//
//	iteam := []string{"cred1", "cred1", "cred2", "cred1"}
//	removeIteam := []string{"cred1"}
//	for index, val := range iteam {
//		misc.DumpObject()
//		if misc.StringInSlice(val, removeIteam) {
//			mapDat[index] = val
//		}
//		for _, val := range mapDat {
//			misc.RemoveStringFromSlice(val, iteam)
//		}
//
//	}
//}
