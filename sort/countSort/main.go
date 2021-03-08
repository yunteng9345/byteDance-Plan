package main

func main() {
	countingSort([]int{3, 5, 6, 2, 2, 6}, 6)
}

func countingSort(arr []int, max int) []int {
	bucketLen := max + 1
	bucket := make([]int, bucketLen)
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]] += 1
	}

	sortIndex := 0
	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			arr[sortIndex] = j
			sortIndex += 1
			bucket[j] -= 1
		}
	}

	return arr
}
