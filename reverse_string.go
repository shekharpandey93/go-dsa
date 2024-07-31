package main

import "fmt"

func main() {
	reverseString("shekhar pandey")

}

func reverseString(str string) string {
	chars := []rune(str)

	for i := 0; i < len(chars)/2; i++ {

		chars[i], chars[len(chars) -i -1] = chars[len(chars) -i -1], chars[i]
	}
	fmt.Println(string(chars))
	return string(chars)
}
