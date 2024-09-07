package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "waterbottle"
	s2 := "erbottlewat"

	if isRotation(s1, s2) {
		fmt.Println("Yes, s2 is a rotation of s1")
	} else {
		fmt.Println("No, s2 is not a rotation of s1")
	}
}

func isRotation(str1 string, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}

	return strings.Contains(str1+str1, str2)
}
