package main

import "slices"

func chocolateDistribution(arr []int, N int) int {
	if len(arr) < N {
		return -1
	}
	slices.Sort(arr)
	resp := arr[N-1] - arr[0]

	for i := 1; (N + i - 1) < N; i++ {
		resp = min(resp, arr[N+i-1]-arr[i])
	}

	return resp
}
