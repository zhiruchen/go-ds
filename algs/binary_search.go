package algs

// BinarySearch 二分查找 返回val的位置
func BinarySearch(array []int32, val int32) int {
	start, end := 0, len(array)-1

	for start <= end {
		mid := (start + end) / 2
		if array[mid] < val {
			start = mid + 1
		} else if array[mid] > val {
			end = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
