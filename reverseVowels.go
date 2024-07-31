package main

import (
	"strings"
)

func main()  {
	reverseVowels("hello")
}

func reverseVowels(s string) {
	str := []rune(s)
	i, j := 0,len(str)-1
	vowels := "aeiouAEIOU"
	for i < j{

		for i < j && !strings.Contains(vowels, string(str[i])) {
			i++
		}

		for i < j && !strings.Contains(vowels, string(str[j])) {
			j--
		}

		if i < j {
			str[i], str[j] = str[j], str[i]
			i++
			j--
		}
	}
}

