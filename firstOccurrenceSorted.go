package main

func main() {

}
func firstOccurrenceinSorted(arr []int, x int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if x > arr[mid] {
			high = mid - 1
		} else if arr[mid] > x {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != arr[mid] {
				return mid
			} else {
				high = mid - 1
			}
		}

	}
	return -1
}
func firstOccurrenceinSortedRecursive(arr []int, low, high, x int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if x > arr[mid] {
		return firstOccurrenceinSortedRecursive(arr, low, mid+1, x)
	} else if x < arr[mid] {
		return firstOccurrenceinSortedRecursive(arr, low, mid-1, x)
	} else {
		if mid == 0 || arr[mid-1] != arr[mid] {
			return mid
		} else {
			return firstOccurrenceinSortedRecursive(arr, low, mid-1, x)
		}
	}
}
