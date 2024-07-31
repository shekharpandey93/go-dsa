package main

import (
	"fmt"
	"math"
)

//[2,3,1,1,4]
//[3,2,1,0,4]
//[3,1,2,0,4]
func main()  {
	arr := []int{1, 3, 6, 1, 0, 9}
	a := getMinimumJump(arr)
	fmt.Println(a)
}

func getMinimumJump(arr []int) int {

	jump := make([]int, len(arr))
	var i int, var j int
	if (len(arr) == 0 || arr[0] == 0) {
		
	}

	if len(arr) <=1 {
		return 0
	}

	currMaxReach := arr[0]
	stepCount := arr[0]
	jump := 0

	for start := 1; start < len(arr) -1; start = start +1 {
		currMaxReach = int(math.Max(float64(currMaxReach), float64(start+ arr[start])))
		stepCount = stepCount + 1
		if stepCount == 0 {
			jump = jump + 1
			stepCount = currMaxReach - start
		}
	}
 	return jump + 1
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

