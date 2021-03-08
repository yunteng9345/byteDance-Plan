package main

func main() {
	BubbleSort11([]int{1, 5, 7, 8, 4, 2})
}

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

// practice
func BubbleSort11(arr []int) []int {
	return nil
}
