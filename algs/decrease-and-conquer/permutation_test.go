package decrease_and_conquer

import (
	"reflect"
	"testing"
)

func TestGenPermutation(t *testing.T) {
	testCases := []struct {
		input  []int32
		result [][]int32
	}{
		{
			input: []int32{1, 2, 3},
			result: [][]int32{
				{2, 3, 1},
				{3, 2, 1},
				{3, 1, 2},
				{1, 3, 2},
				{2, 1, 3},
				{1, 2, 3},
			},
		},
		{
			input:  []int32{1, 2},
			result: [][]int32{{2, 1}, {1, 2}},
		},
	}

	for _, cc := range testCases {
		var res = &[][]int32{}
		GenPermutation(cc.input, len(cc.input), res)

		if !reflect.DeepEqual(*res, cc.result) {
			t.Errorf("expect: %v, get: %v\n", cc.result, *res)
		}
	}
}
