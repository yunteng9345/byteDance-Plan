package main

func main() {
	qucikSort([]int{200, 3, 2, 1, 5, 7, 435, 2, 245, 435, 23, 76, 23, 2, 0})
}

func qucikSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

func sort(arr []int, left int, right int) {
	if left < right {
		key := arr[(left+right)/2]
		i := left
		j := right
		for {
			for arr[i] < arr[key] {
				i++
			}
			for arr[j] > arr[key] {
				j--
			}
			if i >= j {
				break
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
		sort(arr,left,i-1)
		sort(arr,j+1,right)
	}
}
