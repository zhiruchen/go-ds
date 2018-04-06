package dynamic_programming

import "testing"

func TestCoinRow(t *testing.T) {
	cases := []struct {
		coins     []int32
		maxAmount int32
	}{
		{
			coins:     []int32{0, 1, 2, 100},
			maxAmount: 101,
		},
		{
			coins:     []int32{0, 6, 8, 12, 199},
			maxAmount: 207,
		},
		{
			coins:     []int32{0, 5, 1, 2, 10, 6, 2},
			maxAmount: 17,
		},
	}

	for _, ca := range cases {
		result := CoinRow(ca.coins)
		if result != ca.maxAmount {
			t.Errorf("expect: %d, get: %d\n", ca.maxAmount, result)
		}
	}
}
