package main

import "fmt"

func main()  {
	//fmt.Println("Number is prime", isPrimeNeil(101))
	//fmt.Println("Number is prime", isPrime(10))
	getListoFPrimeNumber(10)
}

func isPrimeNeil(num int) bool {
	for idx := 2; idx <= num; idx++ {
		if num%idx == 0 {
			return false
		}
	}
	return true
}

func isPrime(num int) bool {
	if num == 1 {
		return false
	}
	if num == 2 ||  num == 3{
		fmt.Println(num)
		return true
	}
	if num % 2 == 0 || num % 3 == 0 {
		return false
	}
	for idx := 5; idx*idx <= num; idx = idx+6 {
		if num%idx == 0 || num%(idx+2) == 0{
			return false
		}
		fmt.Println(idx)
	}
	return true
}

// Sieve of Eratorthener algo
func getListoFPrimeNumber(num int)  {

	arr := make([]bool, num+1)
	fmt.Println(arr)

	//if num >= 4 {
	//	fmt.Println(2)
	//	fmt.Println(3)
	//}
	//for idx := 5; idx <= num; idx++ {
	//	if isPrime(idx) {
	//		fmt.Println(idx)
	//	}
	//}

	// mapData := make(map[int]bool)
	//mapData[2] = true
	//mapData[3] = true
	//for idx := 4; idx <= num; idx++ {
	//	mapData[idx] = true
	//}
	//for key, val := range mapData {
	//	if key%key != 0 {
	//		mapData[key] = false
	//	}
	//	fmt.Println(key, val)
	//}
}

