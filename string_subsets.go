package main

import "fmt"

func main() {
	subset("sadbutsad", "", 0)
}

func subset(str string, current string, idx int) {
	//var resp []string
	if idx == len(str) {
		//resp = append(resp, current)
		fmt.Println("resp", current)
		return
	}
	subset(str, current, idx+1)
	subset(str, current+string(str[idx]), idx+1)
}
