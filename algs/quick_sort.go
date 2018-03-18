package algs

func QuickSort(array []int32) {
	quickSort(array, 0, len(array)-1)
}

func quickSort(array []int32, low, high int) {
	if low >= high {
		return
	}

	j := partition(array, low, high)
	quickSort(array, low, j-1)
	quickSort(array, j+1, high)
}

func partition(array []int32, low, high int) int {
	pivot := array[low]
	i, j := low, high+1

	for {
		i++
		for array[i] < pivot {
			if i == high {
				break
			}
			i++
		}

		j--
		for pivot < array[j] {
			if j == low {
				break
			}
			j--
		}

		if i >= j {
			break
		}

		array[i], array[j] = array[j], array[i]
	}

	array[low], array[j] = array[j], array[low]
	return j
}

func SelectKthSmallElement(array []int32, k int) int32 {
	if  k < 0 || k >= len(array) {
		panic("k is invalid")
	}

	low, high := 0, len(array)-1

	for high > low {
		p := partition(array, low, high)

		if p > k {
			high = p - 1
		} else if p < k {
			low = p + 1
		} else {
			return array[p]
		}
 	}

 	return array[low]
}