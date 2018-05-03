package dynamic_programming

// MaxProductOfSubArray 获取连续乘积子串的最大值
func MaxProductOfSubArray(array []int32) int32 {
	if len(array) == 0 {
		return -1
	}

	result, maxVal, minVal := array[0], array[0], array[0]

	for j := 1; j < len(array); j++ {
		tmpMax := maxVal * array[j]
		tmpMin := minVal * array[j]

		maxVal = max(max(tmpMax, tmpMin), array[j])
		minVal = min(min(tmpMax, tmpMin), array[j])

		result = max(result, maxVal)
	}

	return result
}

func min(a, b int32) int32 {
	if a <= b {
		return a
	}

	return b
}
