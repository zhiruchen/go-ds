package advanced

import "testing"

func TestRecurReverseString(t *testing.T) {
	testCases := []struct{
		input []string
		output []string
	}{
		{
			[]string{"1", "2", "3"},
			[]string{"3", "2", "1"},
		},
		{
			[]string{"a","b","c"},
			[]string{"c","b","a"},
		},

		{
			[]string{"x","y","z","123"},
			[]string{"123", "z","y","x"},
		},
	}

	for _, c := range testCases {
		c.input = RecurReverseString(c.input)
		if !listEqual(c.input,c.output) {
			t.Errorf("reverse str list, expect: %v, get: %v\n", c.output, c.input)
		}
	}
}

func listEqual(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}

	return true
}