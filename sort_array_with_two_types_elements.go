package main

import "fmt"

func main() {
	arr := []int{-12, 18, -10, 15}
	val := SortArrayWithTwoTypesElements(arr)
	fmt.Println(val)
}
func SortArrayWithTwoTypesElements(nums []int) []int {
	i := -1
	j := len(nums)
	for {
		for {
			i++
			if nums[i] < 0 {
				continue
			} else {
				break
			}
		}

		for {
			j--
			if nums[j] >= 0 {
				continue
			} else {
				break
			}
		}

		if i >= j {
			return nums
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
