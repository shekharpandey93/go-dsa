package main

import "fmt"

func main()  {
	arr := []int{10, 5, 4, 1, 3, 5, 6} //output 4
	index := getRepeatingElement(arr)
	fmt.Println("===", index)
	if index != -1 {
		fmt.Printf("First repeating element is %v and min index is %v", arr[index], index)
	} else {
		fmt.Println("Not found")
	}

}

func getRepeatingElement(arr []int) int  {
	charEle := make(map[int]int)
	//minIndex := -1
	var repeatVal = -1
	for idx := 0; idx < len(arr); idx++ {
		if _, ok := charEle[arr[idx]]; ok{
			repeatVal = arr[idx]
			break
		} else {
			charEle[arr[idx]] = idx
		}
	}
	if repeatVal == -1 {
		return repeatVal
	} else {
		return charEle[repeatVal]
	}
}

// [latest index amd value count]
