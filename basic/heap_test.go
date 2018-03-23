package basic

import (
	"reflect"
	"testing"
)

func TestBinaryHeap_DeleteRoot(t *testing.T) {
	testCases := []struct {
		array  []Comparer
		result Comparer
		vals   []Comparer
	}{
		{
			array:  []Comparer{testVal(0), testVal(89), testVal(12), testVal(101), testVal(3), testVal(66)},
			result: testVal(101),
			vals:   []Comparer{testVal(89), testVal(12), testVal(66), testVal(0), testVal(3)},
		},
		{
			array:  []Comparer{testVal(66), testVal(89), testVal(101)},
			result: testVal(101),
			vals:   []Comparer{testVal(89), testVal(66)},
		},
	}

	for _, cc := range testCases {
		heap := NewBinaryHeap()

		for _, e := range cc.array {
			heap.Insert(e)
		}

		root := heap.DeleteRoot()
		if !root.Equal(cc.result) {
			t.Errorf("expect: %v, get: %v\n", cc.result, root)
		}

		vals := heap.LevelOrder()
		if !reflect.DeepEqual(vals, cc.vals) {

		}
	}
}

func TestBinaryHeap_Sort(t *testing.T) {
	testCases := []struct {
		array  []Comparer
		result []Comparer
	}{
		{
			array:  []Comparer{testVal(0), testVal(89), testVal(12), testVal(101), testVal(3), testVal(66)},
			result: []Comparer{testVal(101), testVal(89), testVal(66), testVal(12), testVal(3), testVal(0)},
		},
		{
			array:  []Comparer{testVal(66), testVal(89), testVal(101)},
			result: []Comparer{testVal(101), testVal(89), testVal(66)},
		},
	}

	for _, cc := range testCases {
		heap := NewBinaryHeap()

		for _, e := range cc.array {
			heap.Insert(e)
		}

		vals := heap.Sort()
		if !reflect.DeepEqual(vals, cc.result) {
			t.Errorf("expect: %v, get: %v\n", cc.result, vals)
		}
	}
}

func TestBinaryHeap_Insert(t *testing.T) {
	testCases := []struct {
		array  []Comparer
		result []Comparer
	}{
		{
			array:  []Comparer{testVal(0), testVal(89), testVal(12), testVal(101), testVal(3), testVal(66)},
			result: []Comparer{testVal(101), testVal(89), testVal(66), testVal(0), testVal(3), testVal(12)},
		},
		{
			array:  []Comparer{testVal(66), testVal(89), testVal(101)},
			result: []Comparer{testVal(101), testVal(66), testVal(89)},
		},
	}

	for _, cc := range testCases {
		heap := NewBinaryHeap()

		for _, e := range cc.array {
			heap.Insert(e)
		}

		vals := heap.LevelOrder()
		if !reflect.DeepEqual(vals, cc.result) {
			t.Errorf("expect: %v, get: %v\n", cc.result, vals)
		}
	}
}
