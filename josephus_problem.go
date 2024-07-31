package main

func main()  {
	josephusProblem(5, 3)
}

func josephusProblem(arr, k int) int  {
	if arr == 1 {
		return 0
	}
	return (josephusProblem(arr-1, k) + k)%arr
}
