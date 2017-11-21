package advanced

import "testing"

func TestTST_Contains(t *testing.T) {
	testCases := []struct {
		input  []string
		target string
		result bool
	}{
		{
			[]string{"three", "simple", "words"},
			"simple",
			true,
		},
		{
			[]string{"three", "simple", "words"},
			"time",
			false,
		},
	}

	for _, c := range testCases {
		tree := NewTST()

		for _, s := range c.input {
			tree.Put(s)
		}

		result := tree.Contains(c.target)
		if result != c.result {
			t.Errorf("expec: %v, get: %v\n", result, c.result)
		}
	}

}
