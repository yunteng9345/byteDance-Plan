package main

// 快速排序思路：
// 1. 从数列中取一个数作为基准数
// 2. 分区过程，将比这个数大的放在基准数的右边，小于或者等于的放左边
// 3. 左右重复第二步，直到各区间只有一个数
// 最坏情况下是O(N²)，平均的时间复杂度是O(N*lgN)。

func main() {
	QuickSort([]int{200, 3, 2, 1, 5, 7, 435, 2, 245, 435, 23, 76, 23, 2, 0})
}

func QuickSort(values []int) {
	sort(values, 0, len(values)-1)
}

// 快速排序:方式二
func sort1(sortArray []int, left, right int) {
	if left < right {
		key := sortArray[(left+right)/2]
		i := left
		j := right
		for {
			for sortArray[i] < key {
				i++
			}
			for sortArray[j] > key {
				j--
			}
			if i >= j {
				break
			}
			sortArray[i], sortArray[j] = sortArray[j], sortArray[i]
		}
		sort1(sortArray, left, i-1)
		sort1(sortArray, j+1, right)
	}
}

