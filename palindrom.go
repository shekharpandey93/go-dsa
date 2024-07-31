package main

import "fmt"

func main()  {
	//str := "日本語本日" // true
	// str := "日本語語本" // false
	 str := "abbccbba" // false
	// str := "日本語語本日" // true
	// str := "madam" // true
	 // str := "maddam" // true
	fmt.Println("Given string is palingdrom %v",isPalindrom(str))
	fmt.Println("Given Number is palingdrom %v",isPalindromNumber(343))
}

func isPalindrom(str string) bool  {
	charString := []rune(str)
	if len(str) == 1{
		return true
	}
	for i := 0; i < len(charString)/2; i++ {
		if (charString[i] != charString[len(charString) - i - 1]) {
			break
		}
		return true
	}
	return false
}

func isPalindromNumber(num int) bool  {
   var rev = 0
	temp := num
	for temp !=0 {
		rev = rev*10+temp%10
		temp = temp/10
		fmt.Println("rev", rev, temp)
   }
	return rev == num
}
