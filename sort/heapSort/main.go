package main

func main() {
	heapSort([]int{3, 2, 3, 1, 2, 4, 5, 5, 6})
}

//func heapSort(arr []int) []int {
//	arrLen := len(arr)
//	buildMaxHeap(arr, arrLen)
//	for i := arrLen - 1; i >= 0; i-- {
//		swap(arr, 0, i)
//		arrLen -= 1
//		heapify(arr, 0, arrLen)
//	}
//	return arr
//}
//
//func buildMaxHeap(arr []int, arrLen int) {
//	for i := arrLen / 2; i >= 0; i-- {
//		heapify(arr, i, arrLen)
//	}
//}
//
//func heapify(arr []int, i int, arrLen int) {
//	left := 2*i + 1
//	right := 2*i + 2
//	largest := i
//	if left < arrLen && arr[left] > arr[largest] {
//		largest = left
//	}
//	if right < arrLen && arr[right] > arr[largest] {
//		largest = right
//	}
//	if largest != i {
//		swap(arr, i, largest)
//		heapify(arr, largest, arrLen)
//	}
//}
//
//func swap(arr []int, i int, j int) {
//	arr[i], arr[j] = arr[j], arr[i]
//}
