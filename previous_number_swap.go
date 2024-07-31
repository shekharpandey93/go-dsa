package main

import (
	"fmt"
	"strings"
)


func main() {
	num := "12435"
	s := "123"
	ps := &s
	b := []byte(*ps)
	fmt.Println("bye",b)
	pb := &b
	s += "4"
	*ps += "5"
	b[1] = 0
	fmt.Print("", *ps)
	fmt.Println("====")

	for {
		fmt.Println("sp")
	}
	fmt.Print(string(*pb))
	previousNumberSwap(num)
}

func previousNumberSwap(n string) string{
	arr := strings.Split(n,"")
	for i := len(arr) -1 ; i >0 ; i--{
		if arr[i]< arr[i-1] {
				temp := arr[i]
				arr[i] = arr[i-1]
				arr[i-1] = temp
			}
		}
	return strings.Join(arr, "")
}
