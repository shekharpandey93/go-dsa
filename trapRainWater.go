package main

import "fmt"

func main() {
	val := getRainWater([]int{5, 0, 6, 2, 3})
	fmt.Println("val", val)
}

func getRainWater(arr []int) int {
	resp := 0
	lMax := make([]int, len(arr))
	for i := 1; i < len(arr); i++ {
		lMax[i] = max(arr[i], lMax[i-1])
	}

	rMax := make([]int, len(arr))
	//rMax[len(arr)-1] = arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		rMax[i] = max(arr[i], rMax[i+1])
	}
	for i := 1; i < len(arr)-1; i++ {

		resp = resp + (getMin(lMax[i], rMax[i]) - arr[i])
		resp += max(0, min(lMax[i], rMax[i])-arr[i])
	}
	return resp
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getRainWater1(height []int) int {
	firstIDX := 0
	lastIDX := len(height) - 1
	maxl := height[firstIDX]
	maxr := height[lastIDX]
	totalWater := 0
	for firstIDX <= lastIDX {
		if maxl > maxr {
			if height[lastIDX] > maxr {
				maxr = height[lastIDX]
			} else if maxr-height[lastIDX] > 0 {
				totalWater += (maxr - height[lastIDX])
			}
			lastIDX--
		} else {
			if height[firstIDX] > maxl {
				maxl = height[firstIDX]
			} else if maxl-height[firstIDX] > 0 {
				totalWater += (maxl - height[firstIDX])
			}
			firstIDX++
		}
	}
	return totalWater
}
