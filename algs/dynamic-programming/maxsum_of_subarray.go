package dynamic_programming

import "math"

// MaxSumOfSubArray
// given a one dimensional array that may contain both positive and negative integers
// find the sum of contiguous subarray of numbers which has the largest sum.
func MaxSumOfSubArray(array []int32) int32 {
	if len(array) == 0 {
		return math.MinInt32
	}

	var maxSum = array[0]
	var sum int32 = 0
	for _, v := range array {
		sum = max(sum+v, v)
		maxSum = max(maxSum, sum)
	}

	return maxSum
}
