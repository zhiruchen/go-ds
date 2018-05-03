package dynamic_programming

import "testing"

func TestMaxProductOfSubArray(t *testing.T) {
	cases := []struct {
		array  []int32
		result int32
	}{
		{
			array:  []int32{10, 2, 3, 0, -1, 9, 100},
			result: 900,
		},
		{
			array:  []int32{10, 2, 3, 0, 100},
			result: 100,
		},
		{
			array:  []int32{100},
			result: 100,
		},

		{
			array:  []int32{0, 1, 2000, 4, 5, 0, 9, 600},
			result: 40000,
		},
		{
			array:  []int32{1, 20, 4, 5, 0, 9, 600},
			result: 5400,
		},
		{
			array:  []int32{1, 20, 4, 5, 0, 600},
			result: 600,
		},
	}

	for _, cc := range cases {
		result := MaxProductOfSubArray(cc.array)
		if result != cc.result {
			t.Errorf("expect: %d, get: %d\n", cc.result, result)
		}
	}
}
