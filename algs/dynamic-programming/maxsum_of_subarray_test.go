package dynamic_programming

import "testing"

func TestMaxSumOfArray(t *testing.T) {
	cases := []struct {
		input  []int32
		maxSum int32
	}{
		{
			input:  []int32{1, -1, 0, 1},
			maxSum: 1,
		},
		{
			input:  []int32{1, 9, -10, 0, 8, 1},
			maxSum: 10,
		},
	}

	for _, c := range cases {
		result := MaxSumOfSubArray(c.input)
		if result != c.maxSum {
			t.Errorf("expect: %v,get: %v\n", c.maxSum, result)
		}
	}
}
