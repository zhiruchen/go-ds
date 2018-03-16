package basic

import "testing"

func TestLinkedList_InsertToFront(t *testing.T) {

	testCases := []struct {
		data   []interface{}
		result string
	}{
		{
			data:   []interface{}{"1", "2", "3", "4"},
			result: "4->3->2->1",
		},
		{
			data:   []interface{}{},
			result: "",
		},
		{
			data:   []interface{}{"1"},
			result: "1",
		},
		{
			data:   []interface{}{1, 2, 3, 4},
			result: "4->3->2->1",
		},
	}

	for _, cc := range testCases {
		ll := NewLinkedList()
		for _, dt := range cc.data {
			ll.InsertToFront(dt)
		}

		dataStr := ll.String()
		if cc.result != dataStr {
			t.Errorf("Expected %s, get %s\n", cc.result, dataStr)
		}
	}
}

func TestLinkedList_AppendToTail(t *testing.T) {
	testCases := []struct {
		data   []interface{}
		result string
	}{
		{
			data:   []interface{}{"1", "2", "3", "4"},
			result: "1->2->3->4",
		},
		{
			data:   []interface{}{},
			result: "",
		},
		{
			data:   []interface{}{"1"},
			result: "1",
		},
		{
			data:   []interface{}{1, 2, 3, 4},
			result: "1->2->3->4",
		},
	}

	for _, cc := range testCases {
		ll := NewLinkedList()
		for _, dt := range cc.data {
			ll.AppendToTail(dt)
		}

		dataStr := ll.String()
		if cc.result != dataStr {
			t.Errorf("Expected %s, get %s\n", cc.result, dataStr)
		}
	}
}

func TestLinkedList_DeleteNode(t *testing.T) {
	testCases := []struct {
		data          []interface{}
		deleteElement interface{}
		result        string
	}{
		{
			data:          []interface{}{"1", "2", "3", "4"},
			deleteElement: "1",
			result:        "2->3->4",
		},
		{
			data:          []interface{}{},
			deleteElement: "x",
			result:        "",
		},
		{
			data:          []interface{}{"1"},
			deleteElement: "1",
			result:        "",
		},
		{
			data:          []interface{}{1, 2, 3, 4},
			deleteElement: 4,
			result:        "1->2->3",
		},
		{
			data:          []interface{}{1, 2, 3, 4, 5},
			deleteElement: 3,
			result:        "1->2->4->5",
		},
	}

	for _, cc := range testCases {
		ll := NewLinkedList()
		for _, dt := range cc.data {
			ll.AppendToTail(dt)
		}

		ll.DeleteNode(cc.deleteElement)

		dataStr := ll.String()
		if cc.result != dataStr {
			t.Errorf("Expected %s, get %s\n", cc.result, dataStr)
		}
	}
}
