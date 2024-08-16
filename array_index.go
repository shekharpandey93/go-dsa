package main

import (
	"fmt"
)

// [2,3,1,1,4]
// [3,2,1,0,4]
// [3,1,2,0,4]
func main() {
	arr := []int{3, 1, 2, 0, 4}
	a := getMinimumJump(arr)
	fmt.Println(a)
}

func getMinimumJump(arr []int) bool {

	lastIdx := len(arr) - 1
	for i := len(arr) - 2; i >= 0; i-- {
		if i+arr[i] >= lastIdx {
			lastIdx = i
		}
	}
	return lastIdx == 0
}

//function minJumps(arr , n)
//
//{
//// jumps[n-1] will hold the
//var jumps = Array.from({length: n}, (_, i) => 0);;
//// result
//var i, j;
//
//// if first element is 0,
//if (n == 0 || arr[0] == 0)
//return Number.MAX_VALUE;
//// end cannot be reached
//
//jumps[0] = 0;
//
//// Find the minimum number of jumps to reach arr[i]
//// from arr[0], and assign this value to jumps[i]
//for (i = 1; i < n; i++) {
//jumps[i] = Number.MAX_VALUE;
//for (j = 0; j < i; j++) {
//if (i <= j + arr[j]
//&& jumps[j]
//!= Number.MAX_VALUE) {
//jumps[i] = Math.min(jumps[i], jumps[j] + 1);
//break;
//}
//}
//}
//return jumps[n - 1];
//}
//
//// driver program to test above function
//var arr = [ 1, 3, 6, 1, 0, 9 ];
