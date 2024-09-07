package main

import "fmt"

// [30][30]board;
func main() {
	board := [30]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	dataMap := make(map[int]bool)

	for jdx := 0; jdx < len(board); jdx = 0 {
		for idx := 0; idx < len(board); idx = idx + 1 {
			if _, ok := dataMap[idx]; !ok {
				dataMap[idx] = false
			} else {
				dataMap[idx] = !dataMap[idx]
			}
		}
	}
	fmt.Println(dataMap)
	//for key := range dataMap {
	//	fmt.Println("dataMap",dataMap[key])
	//}
}
