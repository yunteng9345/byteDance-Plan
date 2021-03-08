package main

func main() {
	shellSort([]int{4, 5, 2, 6, 7, 3, 1})
}

func shellSort(arr []int) []int {
	length := len(arr)
	gap := 1
	for gap < gap/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
	}
	return arr
}

func shellSort1(arr []int) []int {
	return nil
}
