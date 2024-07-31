package main

import "fmt"

func main()  {
	//val := getLCMNeil(12, 15)
	val := getLCM(3, 7)
	fmt.Println("val",val)
}

func getLCMNeil(a int, b int)  int{

	max := a
	if b > a {
		max = b
	}

	for {
		if max % a == 0 &&  max % b == 0 {
			return max
		}
			max ++;
	}
	return max
}
// Euclidean formula
// formula a * b = gcd(a,b) * lcm(a,b)
// LCM (a*b)/gcd(a,b)
func getLCM(a int, b int)  int{


	//max := a
	//diff := a - b
	//if b > a {
	//	//max = b
	//	diff = b -a
	//}
	//return a/diff * b/diff * diff

	return (a*b)/getGcd(a,b)

}

func getGcd(a int, b int) int  {
	if b == 0 {
		return a
	}
	return getGcd(b, a%b)
}
