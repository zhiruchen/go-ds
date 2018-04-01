package dynamic_programming

// CoinRow 找出最大数量的钱数从coins列表中，但不能使用相邻的两个coin
func CoinRow(coins []int32) int32 {

	// 这个问题可以被分成两组，第一组中coin里不包含最后一个coin, 第二组包含最后一个coin
	// F(n) = max(Cn + F(n-2), F(n-1)), F(0) = 0, F(1) = C(1)
	// F(2) = max(C2 + F(0), F(1))
	// F(3) = max(C3 + F(1), F(2))
	// F(4) = max(C4 + F(2), F(3))
	n := len(coins)
	f := make([]int32, n)
	f[0], f[1] = 0, coins[1]

	for i := 2; i <= n-1; i++ {
		f[i] = max(coins[i]+f[i-2], f[i-1])
	}

	return f[n-1]
}

func max(a, b int32) int32 {
	if a >= b {
		return a
	}

	return b
}
