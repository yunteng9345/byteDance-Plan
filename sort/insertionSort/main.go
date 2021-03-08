package main

func main() {
	insertionSort1([]int{5, 2, 3, 1, 7, 8})
}
func insertionSort(arr []int) []int {
	for i := range arr {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex = preIndex - 1
		}
		arr[preIndex+1] = current
	}
	return arr
}

// practice
func insertionSort1(arr []int) []int {
	return nil
}
