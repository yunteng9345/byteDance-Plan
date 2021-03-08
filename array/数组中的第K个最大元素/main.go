package main

import "fmt"

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	bulidMacHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		heapify(nums, 0, heapSize)
	}
	return nums[0]
}

func bulidMacHeap(nums []int, size int) {
	for i := size / 2; i >= 0; i-- {
		heapify(nums, i, size)
	}
}

func heapify(nums []int, i int, size int) {
	l, r, largest := i*2+1, i*2+2, i
	for l < size && nums[l] > largest {
		largest = l
	}
	for r < size && nums[r] > largest {
		largest = r
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapify(nums, largest, size)
	}
}
