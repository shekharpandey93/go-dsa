package main

import "fmt"

func main(){
	//O(n)
	getTrailingZeroForFactorial(10)
	//o(lon n)
	getTrailingZeroForFactorialSpace(10)
}

func getTrailingZeroForFactorial(num int)  {
	fact := 1
	for idx:=2; idx <=num; idx++ {
		fact = fact*idx
	}
	count := 0
	for fact%10 == 0 {
		fact = fact/10
		count++
	}
	fmt.Println("fact",count)
}

func getTrailingZeroForFactorialSpace(num int)  {
	count := 0
	for idx:=5; idx<=num; idx = idx*5 {
		count = count + num/idx
	}
	fmt.Println("fact",count)
	//count := 0
	//for fact%10 == 0 {
	//	fact = fact/10
	//	count++
	//}
}

