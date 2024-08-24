package main

import (
	"fmt"
	"strings"
)

func main() {
	reverseString("shekhar pandey")
	reverseWord("shekhar pandey")

}

func reverseString(str string) string {
	chars := []rune(str)

	for i := 0; i < len(chars)/2; i++ {

		chars[i], chars[len(chars)-i-1] = chars[len(chars)-i-1], chars[i]
	}
	fmt.Println(string(chars))
	return string(chars)
}

func reverseWord(str string) {
	words := strings.Fields(str)

	for idx := 0; idx < len(words)/2; idx++ {
		words[idx], words[len(words)-idx-1] = words[len(words)-idx-1], words[idx]
	}

	//strings.Join(words, " ")
	fmt.Println(strings.Join(words, " "))
}
