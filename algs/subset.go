package algs

import (
	"bytes"
	"math"
	"strconv"
)

// GenSubSetOfNum 生成子集
func GenSubSetOfNum(n int64) [][]int64 {
	var result [][]int64
	var i int64

	number := math.Pow(float64(2), float64(n))
	for i = 1; i <= int64(number)-1; i++ {
		binary := strconv.FormatInt(i, 2)

		binLen := int64(len(binary))
		if binLen < n {
			binary = genZeros(n-binLen) + binary
		}

		var innerResult []int64
		for j, c := range binary {
			if c == '1' {
				innerResult = append(innerResult, int64(j+1))
			}
		}

		result = append(result, innerResult)
	}

	return result
}

func genZeros(n int64) string {
	var buf bytes.Buffer
	var i int64
	for i = 1; i <= n; i++ {
		buf.WriteString("0")
	}

	return buf.String()
}
