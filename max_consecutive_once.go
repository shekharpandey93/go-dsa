package main

func main() {
	arr := []int{1, 1, 0, 1, 1, 1, 1, 0, 0, 1, 1, 1}
	resp := 0
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			count = 0
		} else {
			count++
			resp = max(resp, count)
		}
	}
}
