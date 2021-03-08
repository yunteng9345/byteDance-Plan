package main

import (
	"fmt"
)

func main() {
	fmt.Println()
}

//首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置。

//再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。

//重复第二步，直到所有元素均排序完毕。

func selectSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; i < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

// practice
func selectSort1(arr []int) []int {
	return nil
}
