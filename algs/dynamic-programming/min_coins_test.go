package dynamic_programming

import "testing"

func TestMinNumberOfCoins(t *testing.T) {
	testCases := []struct {
		coins []int32
		s     int32
		num   int32
	}{
		{
			coins: []int32{1, 3, 5},
			s:     1,
			num:   1,
		},
		{
			coins: []int32{1, 3, 5},
			s:     3,
			num:   1,
		},
		{
			coins: []int32{1, 3, 5},
			s:     5,
			num:   1,
		},
		{
			coins: []int32{1, 3, 5},
			s:     4,
			num:   2,
		},
		{
			coins: []int32{1, 3, 5},
			s:     8,
			num:   2,
		},

		{
			coins: []int32{1, 3, 5},
			s:     0,
			num:   0,
		},
		{
			coins: []int32{1, 3, 5},
			s:     11,
			num:   3,
		},
	}

	for _, cc := range testCases {
		num := MinNumberOfCoins(cc.coins, cc.s)
		if num != cc.num {
			t.Errorf("expect num: %d, get: %d\n", cc.num, num)
		}
	}
}
