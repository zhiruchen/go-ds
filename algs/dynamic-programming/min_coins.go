package dynamic_programming

import (
	"math"
)

// MinNumberOfCoins 从coins列表里找出和为s的最少的硬币数
// https://www.topcoder.com/community/data-science/data-science-tutorials/dynamic-programming-from-novice-to-advanced/
// min(s) = min(s-1) + 1
// min(s-1) = min(s-2) + 1
// min(s-2) = min(s-3) + 1
// ...
// min(1) = min(0) + 1
// min(0) = 0
func MinNumberOfCoins(coins []int32, s int32) int32 {
	var i int32
	var minNumberCoins = make([]int32, s+1)
	minNumberCoins[0] = 0
	for i = 1; i <= s; i++ {
		minNumberCoins[i] = math.MaxInt32
	}

	var j int32
	for j = 1; j <= s; j++ {
		for idx := 0; idx < len(coins); idx++ {
			v := coins[idx]
			if v <= j && minNumberCoins[j-v]+1 < minNumberCoins[j] {
				minNumberCoins[j] = minNumberCoins[j-v] + 1
			}
		}
	}

	return minNumberCoins[s]
}
