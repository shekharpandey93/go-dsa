package main

import "fmt"

func main()  {
	//val := gcdOfStrings("ABCABC", "ABC" )
	val := gcdOfStrings("ABABAB", "ABAB" )
	fmt.Println("val",val)
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	len := gcd( len(str1), len(str2) )
	return str1[:len]

}

func gcd( a int, b int ) int {

	for b != 0 {
		a , b = b , a%b
		fmt.Println(a, b)
	}

	return a
}

func getGCD(a int, b int) int {
	if b == 0 {
		return  a
	}
	return getGCD(b, a%b)
}
