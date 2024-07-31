package main

import (
	"fmt"
)

func main()  {
	fmt.Println(isAnagram("abca", "abca" ))
}

func isAnagram(str1, str2 string)  bool{
	char1 := []rune(str1)
	char2 := []rune(str2)
	charEle1 := make(map[string]int)
	charEle2 := make(map[string]int)
	if (len(char1) != len(char2)) {
		return false
	}

		for i:=0; i< len(char1); i++ {
		charEle1[string(char1[i])] = charEle1[string(char1[i])] + 1
		charEle2[string(char2[i])] = charEle2[string(char2[i])] + 1
	}
	for val := range charEle1 {
		if charEle1[val] != charEle2[val] {
			return false
		}
	}
	return true
}
