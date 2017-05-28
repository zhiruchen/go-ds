package basic

func IsStrongNumber(a int64) bool {
	if a == 0 {
		return true
	}

	return sumNl(numParse(a)) == a
}

func numParse(n int64) []int64 {
	nl := []int64{}
	for n > 0 {
		x := n % 10
		nl = append(nl, x)
		n /= 10
	}
	return nl
}

func sumNl(nl []int64) int64 {
	var sum int64 = 0
	for _, n := range nl {
		sum += fact(n)
	}
	return sum
}

func fact(n int64) int64 {
	var x int64 = 1
	for i := 1; int64(i) <= n; i++ {
		x *= int64(i)
	}
	return x
}
