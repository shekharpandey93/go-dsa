package main

func main() {
	mergesort([]int{10, 5, 30, 17, 6}, 0, 4)
}

func mergesort(arr []int, l, r int) {
	if l < r {
		mid := l + (r-l)/2

		mergesort(arr, l, mid)
		mergesort(arr, mid+1, r)
		merge(arr, l, mid, r)

	}
}

func merge(arr []int, low, mid, high int) {
	n1 := mid - low + 1
	n2 := high - mid
	left := make([]int, n1)
	right := make([]int, n2)
	for i := 0; i < n1; i++ {
		left[i] = arr[low+i]
	}
	for i := 0; i < n2; i++ {
		right[i] = arr[mid+i+1]
	}
	i, j, k := 0, 0, low
	for i < n1 && j < n2 {
		if left[i] <= right[j] {
			arr[k] = left[i]
			k++
			i++
		} else {
			arr[k] = right[j]
			k++
			j++
		}
	}
	for i < n1 {
		i++
		k++
	}
	for j < n2 {
		j++
		k++
	}
}
