package main

import (
	"fmt"
	"strings"
)

func main()  {
	str := "GeeksforGeeks A computer science portal for geeks portal portal"
	searchString :=  "portal"
	fmt.Println(getStringCount(str, searchString))
}

func getStringCount(str, searchString string) (count int) {
	splitedWords := strings.Split(str, " ")
	for i := 0; i < len(splitedWords); i++ {
		if splitedWords[i] == searchString {
			count += 1
		}
	}
	return count
}
