package decrease_and_conquer

func GenPermutation(array []int32, n int, result *[][]int32) {
	if n == 1 {
		tmpArray := make([]int32, len(array))
		copy(tmpArray, array)
		*result = append(*result, tmpArray)
		return
	}

	for i := 0; i < n; i++ {
		swap(array, i, n-1)
		GenPermutation(array, n-1, result)
		swap(array, i, n-1)
	}
}

func swap(a []int32, x, y int) {
	if x == y {
		return
	}

	a[x], a[y] = a[y], a[x]
}
