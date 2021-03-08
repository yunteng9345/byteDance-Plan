package main
//
//func main() {
//	sortArray([]int{3, 2, 3, 1, 2, 4, 5, 5, 6})
//}
//
//func sortArray(nums []int) []int {
//	if len(nums) < 2 {
//		return nums
//	}
//	i := len(nums) / 2
//	left := sortArray(nums[0:i])
//	right := sortArray(nums[i:])
//
//	return merge(left, right)
//}
//
//func merge(l, r []int) []int {
//	i, j := 0, 0
//	m, n := len(l), len(r)
//	var res []int
//	for i < m && j < n {
//		if l[i] < r[j] {
//			res = append(res, l[i])
//			i++
//		} else {
//			res = append(res, r[j])
//			j++
//		}
//	}
//	res = append(res, l[i:]...)
//	res = append(res, r[j:]...)
//
//	return res
//}
