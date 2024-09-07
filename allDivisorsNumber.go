package main

import "fmt"

func main() {
	getAllDivisorNumber(100)
	//getAllDivisorNumberSorted(15)
}

func getAllDivisorNumber(num int) {
	for idx := 1; idx*idx <= num; idx++ {
		if num%idx == 0 {
			fmt.Println(idx)
			if idx != num/idx {
				fmt.Println(num / idx)
			}
		}
	}
}

func getAllDivisorNumberSorted(num int) {
	var idx int
	for idx = 1; idx*idx <= num; idx++ {
		if num%idx == 0 {
			fmt.Println(idx)
		}
	}
	for ; idx >= 1; idx-- {
		if num%idx == 0 {
			fmt.Println("hi", num/idx)
		}
	}
}
