package main

import "fmt"

//[2,2,1,1,5]
//[4,1,2,1,2]
//[1]

func main()  {
	arr := []int{2,2,1,1,5}
	nonRepeatVal := nonDuplicateVal(arr)
	fmt.Println("nonRepeatVal",nonRepeatVal)


	arr1 := []int{4,1,2,1,2}
	val1 := nonDuplicateVal(arr1)
	fmt.Println("val",val1)


	arr2 := []int{1}
	val2 := nonDuplicateVal(arr2)
	fmt.Println("val2",val2)

}

func nonDuplicateVal(arr []int) (num int) {
	// map
	mapData := make(map[int]int)
	for _, val := range arr {
		 	mapData[val] = mapData[val] + 1

	}

	for key, _ := range mapData {
		if mapData[key] == 1 {
			num =  key
			break
		}
	}
	return
}


