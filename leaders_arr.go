package main

import "fmt"

func main() {
	getLeaders([]int{7, 10, 4, 10, 6, 5, 2})
}

func getLeaders(arr []int) {
	leader := arr[len(arr)-1]
	var leaderArr []int
	leaderArr = append(leaderArr, leader)
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > leader {
			leader = arr[i]
			leaderArr = append(leaderArr, leader)
		}
	}
	fmt.Println(leaderArr)
}
