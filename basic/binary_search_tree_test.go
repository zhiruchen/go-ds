package basic

import (
	"reflect"
	"testing"
)

type testVal int32

func (v testVal) Equal(v2 Comparer) bool {
	vv := v2.(testVal)
	return v == vv
}

func (v testVal) Less(v2 Comparer) bool {
	vv := v2.(testVal)
	return v <= vv
}

func TestBinarySearchTree_Delete(t *testing.T) {
	testCases := []struct {
		vals   []Comparer
		delete Comparer
		result []Comparer
	}{
		{
			vals:   []Comparer{testVal(32), testVal(36), testVal(6), testVal(7)},
			delete: testVal(7),
			result: []Comparer{testVal(32), testVal(6), testVal(36)},
		},
		{
			vals:   []Comparer{testVal(32)},
			delete: testVal(33),
			result: []Comparer{testVal(32)},
		},
		{
			vals:   []Comparer{testVal(32), testVal(36), testVal(6), testVal(7)},
			delete: testVal(32),
			result: []Comparer{testVal(36), testVal(6), testVal(7)},
		},
	}

	for _, cc := range testCases {
		bst := BinarySearchTree{root: nil}
		for _, val := range cc.vals {
			bst.Insert(val)
		}

		node := bst.Search(cc.delete)
		bst.Delete(node)

		vals, err := bst.LevelTraverse()
		if err != nil {
			t.Errorf("expect: nil, get: %v\n", err)
		}

		if !reflect.DeepEqual(cc.result, vals) {
			t.Errorf("expect: %v, get: %v\n", cc.result, vals)
		}
	}
}

func TestBinarySearchTree_Search(t *testing.T) {
	testCases := []struct {
		vals   []Comparer
		search Comparer
		result Comparer
	}{
		{
			vals:   []Comparer{testVal(32), testVal(36), testVal(6), testVal(7)},
			search: testVal(7),
			result: testVal(7),
		},
		{
			vals:   []Comparer{testVal(32)},
			search: testVal(32),
			result: testVal(32),
		},
		{
			vals:   []Comparer{testVal(32), testVal(36), testVal(6), testVal(7)},
			search: testVal(9),
			result: nil,
		},
	}

	for _, cc := range testCases {
		bst := BinarySearchTree{root: nil}
		for _, val := range cc.vals {
			bst.Insert(val)
		}

		node := bst.Search(cc.search)

		if node == nil {
			if cc.result == nil {
				continue
			} else {
				t.Errorf("expect: %v, get: %v\n", cc.result, node)
			}
		}

		if !node.data.Equal(cc.result) {
			t.Errorf("expect: %v, get: %v\n", cc.result, node.data)
		}
	}
}

func TestBinarySearchTree_Insert(t *testing.T) {
	vals := []Comparer{testVal(32), testVal(36), testVal(6), testVal(7)}

	bst := BinarySearchTree{root: nil}
	for _, val := range vals {
		bst.Insert(val)
	}

	orderVals := []Comparer{testVal(32), testVal(6), testVal(36), testVal(7)}
	levelVals, err := bst.LevelTraverse()
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(levelVals, orderVals) {
		t.Errorf("expect: %v, get: %v\n", orderVals, levelVals)
	}
}
