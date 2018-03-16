package basic

import "testing"

func TestStack_Push(t *testing.T) {
	testCases := []struct{
		elements []interface{}
		result string
		isEmpty bool
	}{
		{
			elements:[]interface{}{},
			result:"",
			isEmpty:true,
		},
		{
			elements:[]interface{}{1,2,3,4},
			result:"4->3->2->1",
			isEmpty:false,
		},
	}

	for _, cc := range testCases {
		stack := NewStack()
		for _, e := range cc.elements {
			stack.Push(e)
		}

		stackStr := stack.String()
		if cc.result != stackStr {
			t.Errorf("expectd %s, get %s\n", cc.result, stackStr)
		}

		empty := stack.IsEmpty()
		if cc.isEmpty != empty {
			t.Errorf("expectd %v, get %v\n", cc.isEmpty, empty)
		}
	}
}

func TestStack_Pop(t *testing.T) {
	testCases := []struct{
		elements []interface{}
		result interface{}
		stackStr string
	}{
		{
			elements:[]interface{}{},
			result: nil,
			stackStr: "",
		},
		{
			elements:[]interface{}{1,2,3,4},
			result:4,
			stackStr: "3->2->1",
		},
		{
			elements:[]interface{}{"1", "2", "3", "4"},
			result:"4",
			stackStr: "3->2->1",
		},
	}

	for _, cc := range testCases {
		stack := NewStack()
		for _, e := range cc.elements {
			stack.Push(e)
		}

		val, err := stack.Pop()
		if len(cc.elements) == 0 {
			if err == nil {
				t.Errorf("Pop should return error when no element on stack\n")
			}
		}

		if cc.result != val {
			t.Errorf("expected: %v, get: %v\n", cc.result, val)
		}

		sstr := stack.String()
		if cc.stackStr != sstr {
			t.Errorf("expected: %s, get: %s\n", cc.stackStr, sstr)
		}
	}
}

func TestStack_Peek(t *testing.T) {
	testCases := []struct{
		elements []interface{}
		result interface{}
	}{
		{
			elements:[]interface{}{},
			result: nil,
		},
		{
			elements:[]interface{}{1,2,3,4},
			result:4,
		},
		{
			elements:[]interface{}{"1", "2", "3", "4"},
			result:"4",
		},
	}

	for _, cc := range testCases {
		stack := NewStack()
		for _, e := range cc.elements {
			stack.Push(e)
		}

		val, err := stack.Peek()
		if len(cc.elements) == 0 {
			if err == nil {
				t.Errorf("Pop should return error when no element on stack\n")
			}
		}

		if cc.result != val {
			t.Errorf("expected: %v, get: %v\n", cc.result, val)
		}
	}
}