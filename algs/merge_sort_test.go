package algs

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		input  []int32
		output []int32
	}{
		{
			input:  []int32{1, 2, 3, 4, 5},
			output: []int32{1, 2, 3, 4, 5},
		},
		{
			input:  []int32{7, 3, 2, 5, 1, 6},
			output: []int32{1, 2, 3, 5, 6, 7},
		},
		{
			input:  []int32{7, 3, 2, 5, 1, 6, 0},
			output: []int32{0, 1, 2, 3, 5, 6, 7},
		},
		{
			input:  []int32{100},
			output: []int32{100},
		},
	}

	for _, cc := range testCases {
		result := MergeSort(cc.input)
		if !reflect.DeepEqual(result, cc.output) {
			t.Errorf("expected: %v, get: %v\n", cc.output, result)
		}
	}
}
