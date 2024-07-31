package main

import (
	"fmt"
	"strings"
)

func main()  {
	str := "The quick brown fox jumps over a lazy dog F."
	fmt.Println(isPangram(str))
}

func isPangram(str string)  bool {
	charEle := make(map[string]bool)
	for _,s := range str {
		if (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z') {
			charEle[strings.ToLower(string(s))] =  true
		}
	}
	if len(charEle) != 26 {
		return false
	}

	return true
}
