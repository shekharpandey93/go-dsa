package main

import "fmt"

// Two Sum II - Input Array Is Sorted
// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
//
// Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
//
// The tests are generated such that there is exactly one solution. You may not use the same element twice.
//
// Your solution must use only constant extra space.
func main() {
	val := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println("val", val)
}

func twoSum(numbers []int, target int) []int {
	//for left, right := 0, len(numbers)-1; ; {
	//	switch {
	//	case numbers[left]+numbers[right] > target:
	//		right--
	//	case numbers[left]+numbers[right] < target:
	//		left++
	//	default:
	//		return []int{left, right}
	//	}
	//}
	resp := make(map[int]int)

	for i, num := range numbers {
		if j, ok := resp[target-num]; ok {
			return []int{i, j}
		}

		resp[num] = i
	}

	return []int{}

}

// var abc []int
// idx:= 0
// jdx:= 1
// for idx <= len(nums)-1 {
//     if jdx <= len(nums)-1 && (nums[idx] + nums[jdx]) == target {
//         abc = append(abc, idx)
//         abc = append(abc, jdx)
//          return abc
//     }
//     if jdx == len(nums) {
//         idx++
//         jdx = idx
//     }
//     jdx++
// }
// return abc

//numbers := make(map[int]int)
//
//for i, num := range nums {
//if j, ok := numbers[target-num]; ok {
//return []int{i, j}
//}
//
//numbers[num] = i
//}
//
//return []int{}

// 	for left, right := 0, len(nums) - 1 ;; {
// 	switch {
// 		case nums[left] + nums[right] > target: right--
// 		case nums[left] + nums[right] < target: left++
// 		default: return []int{left, right}
// 	}
// }
