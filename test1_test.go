package main

import (
	"fmt"
	"testing"
)

func TestDuplicateVal(t *testing.T)  {
	arr := []int{2,2,1,1,5}
	val := NonDuplicateVal(arr)

	if val == 5 {
		fmt.Println("Test Pass 1")
	}

	arr1 := []int{4,1,2,1,2}
	val1 := NonDuplicateVal(arr1)
	if val1 == 4 {
		fmt.Println("Test Pass 2")
	}

	arr2 := []int{1}
	val2 :=  NonDuplicateVal(arr2)

	if val2 == 1 {
		fmt.Println("Test Pass 3")
	}


}
