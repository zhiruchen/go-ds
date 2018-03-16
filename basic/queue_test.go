package basic

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	testCases := []struct {
		input  []interface{}
		output string
	}{
		{
			input:  []interface{}{"A", "B", "C", "D"},
			output: "A->B->C->D",
		},
		{
			input:  []interface{}{1, 2, 3, 4},
			output: "1->2->3->4",
		},
		{
			input:  []interface{}{},
			output: "",
		},
	}

	for _, cc := range testCases {
		q := NewQueue()
		for _, e := range cc.input {
			q.Enqueue(e)
		}

		qStr := q.String()
		if cc.output != qStr {
			t.Errorf("expected: %s, get: %s\n", cc.output, qStr)
		}
	}
}

func TestQueue_Dequeue(t *testing.T) {
	testCases := []struct {
		input  []interface{}
		output interface{}
		qStr   string
	}{
		{
			input:  []interface{}{"A", "B", "C", "D"},
			output: "A",
			qStr:   "B->C->D",
		},
		{
			input:  []interface{}{1, 2, 3, 4},
			output: 1,
			qStr:   "2->3->4",
		},
		{
			input:  []interface{}{},
			output: nil,
			qStr:   "",
		},
	}

	for _, cc := range testCases {
		q := NewQueue()
		for _, e := range cc.input {
			q.Enqueue(e)
		}

		v, err := q.Dequeue()
		if len(cc.input) == 0 {
			if err == nil {
				t.Errorf("should get error when no elements")
			}
		} else {
			if err != nil {
				t.Errorf("should get no error when have elements")
			}
		}

		if v != cc.output {
			t.Errorf("expected: %v, get: %v\n", cc.output, v)
		}

		ss := q.String()
		if cc.qStr != ss {
			t.Errorf("expected: %s, get: %s\n", cc.qStr, ss)
		}
	}
}
