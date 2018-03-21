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

func TestComparer(t *testing.T) {
}

func TestBinarySearchTree_Search(t *testing.T) {

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
