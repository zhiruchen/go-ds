package algs

import "testing"

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		vals []int32
		val  int32
		pos  int
	}{
		{
			vals: []int32{-1, 0, 2, 3, 10},
			val:  -1,
			pos:  0,
		},
		{
			vals: []int32{-1, 0, 2, 3, 10},
			val:  10,
			pos:  4,
		},
		{
			vals: []int32{-1, 0, 2, 3, 10},
			val:  11,
			pos:  -1,
		},
		{
			vals: []int32{-1, 0, 2, 3, 10},
			val:  2,
			pos:  2,
		},
	}

	for _, cc := range testCases {
		result := BinarySearch(cc.vals, cc.val)
		if result != cc.pos {
			t.Errorf("expected: %d, get: %d\n", cc.pos, result)
		}
	}
}
