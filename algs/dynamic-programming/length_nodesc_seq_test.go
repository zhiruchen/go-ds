package dynamic_programming

import "testing"

func TestLengthOfLongestNoDescSeq(t *testing.T) {
	testCases := []struct {
		seq    []int32
		length int32
	}{
		{
			seq:    []int32{5, 3, 4, 8, 6, 7},
			length: 4,
		},
		{
			seq:    []int32{6, 7, 8, 9},
			length: 4,
		},
		{
			seq:    []int32{0, 1, 2, 10, 3, 4, 5, 7, 9},
			length: 8,
		},
		{
			seq:    []int32{0, 1, -1, 2, 1, 6},
			length: 4,
		},
		{
			seq:    []int32{6, 5, 4, 3, 2, 1},
			length: 1,
		},
	}

	for _, cc := range testCases {
		length := LengthOfLongestNoDescSeq(cc.seq)
		if length != cc.length {
			t.Errorf("expect length: %d, get: %d\n", cc.length, length)
		}
	}
}
