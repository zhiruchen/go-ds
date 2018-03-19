package algs

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
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
		QuickSort(cc.input)
		if !reflect.DeepEqual(cc.input, cc.output) {
			t.Errorf("expected: %v, get: %v\n", cc.output, cc.input)
		}
	}
}

func TestSelectKthSmallElement(t *testing.T) {
	testCases := []struct {
		input  []int32
		k      int
		output int32
	}{
		{
			input:  []int32{1, 2, 3, 4, 5},
			k:      3,
			output: 4,
		},
		{
			input:  []int32{7, 3, 2, 5, 1, 6},
			k:      2,
			output: 3,
		},
		{
			input:  []int32{7, 3, 2, 5, 1, 6, 0},
			k:      5,
			output: 6,
		},
		{
			input:  []int32{100},
			k:      0,
			output: 100,
		},
	}

	for _, cc := range testCases {
		result := SelectKthSmallElement(cc.input, cc.k)
		if result != cc.output {
			t.Errorf("expected: %v, get: %v\n", cc.output, result)
		}
	}
}
