package basic

import (
	"reflect"
	"testing"
)

func TestBinaryTree_InOrder(t *testing.T) {
	testCases := []struct {
		node   *BinTreeNode
		result []interface{}
	}{
		{
			node: &BinTreeNode{
				data: 1,
				left: &BinTreeNode{
					data: 2,
					left: &BinTreeNode{
						data: 3,
						left: nil,
						right: &BinTreeNode{
							data:  4,
							left:  nil,
							right: nil,
						},
					},
					right: &BinTreeNode{
						data: 5,
						left: &BinTreeNode{
							data:  6,
							left:  nil,
							right: nil,
						},
						right: &BinTreeNode{
							data:  7,
							left:  nil,
							right: nil,
						},
					},
				},
				right: &BinTreeNode{
					data:  8,
					left:  nil,
					right: nil,
				},
			},
			result: []interface{}{3, 4, 2, 6, 5, 7, 1, 8},
		},
	}

	for _, cc := range testCases {
		binTree := NewBinaryTree(cc.node)
		values := binTree.InOrder()
		if !reflect.DeepEqual(values, cc.result) {
			t.Errorf("Expected %v, get: %v", cc.result, values)
		}
	}
}

func TestBinaryTree_PreOrder(t *testing.T) {
	testCases := []struct {
		node   *BinTreeNode
		result []interface{}
	}{
		{
			node: &BinTreeNode{
				data: 1,
				left: &BinTreeNode{
					data: 2,
					left: &BinTreeNode{
						data: 3,
						left: nil,
						right: &BinTreeNode{
							data:  4,
							left:  nil,
							right: nil,
						},
					},
					right: &BinTreeNode{
						data: 5,
						left: &BinTreeNode{
							data:  6,
							left:  nil,
							right: nil,
						},
						right: &BinTreeNode{
							data:  7,
							left:  nil,
							right: nil,
						},
					},
				},
				right: &BinTreeNode{
					data:  8,
					left:  nil,
					right: nil,
				},
			},
			result: []interface{}{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, cc := range testCases {
		binTree := NewBinaryTree(cc.node)
		values := binTree.PreOrder()
		if !reflect.DeepEqual(values, cc.result) {
			t.Errorf("Expected %v, get: %v", cc.result, values)
		}
	}
}

func TestBinaryTree_PostOrder(t *testing.T) {
	testCases := []struct {
		node   *BinTreeNode
		result []interface{}
	}{
		{
			node: &BinTreeNode{
				data: 1,
				left: &BinTreeNode{
					data: 2,
					left: &BinTreeNode{
						data: 3,
						left: nil,
						right: &BinTreeNode{
							data:  4,
							left:  nil,
							right: nil,
						},
					},
					right: &BinTreeNode{
						data: 5,
						left: &BinTreeNode{
							data:  6,
							left:  nil,
							right: nil,
						},
						right: &BinTreeNode{
							data:  7,
							left:  nil,
							right: nil,
						},
					},
				},
				right: &BinTreeNode{
					data:  8,
					left:  nil,
					right: nil,
				},
			},
			result: []interface{}{4, 3, 6, 7, 5, 2, 8, 1},
		},
	}

	for _, cc := range testCases {
		binTree := NewBinaryTree(cc.node)
		values := binTree.PostOrder()
		if !reflect.DeepEqual(values, cc.result) {
			t.Errorf("Expected %v, get: %v", cc.result, values)
		}
	}
}

func TestBinaryTree_LevelOrder(t *testing.T) {
	testCases := []struct {
		node   *BinTreeNode
		result []interface{}
	}{
		{
			node: &BinTreeNode{
				data: 1,
				left: &BinTreeNode{
					data: 2,
					left: &BinTreeNode{
						data: 3,
						left: nil,
						right: &BinTreeNode{
							data:  4,
							left:  nil,
							right: nil,
						},
					},
					right: &BinTreeNode{
						data: 5,
						left: &BinTreeNode{
							data:  6,
							left:  nil,
							right: nil,
						},
						right: &BinTreeNode{
							data:  7,
							left:  nil,
							right: nil,
						},
					},
				},
				right: &BinTreeNode{
					data:  8,
					left:  nil,
					right: nil,
				},
			},
			result: []interface{}{1, 2, 8, 3, 5, 4, 6, 7},
		},
	}

	for _, cc := range testCases {
		binTree := NewBinaryTree(cc.node)
		values, err := binTree.LevelOrder()
		t.Logf("values: %v\n", values)
		if err != nil {
			t.Errorf("level order should not return error, err: %v", err)
		}

		if !reflect.DeepEqual(values, cc.result) {
			t.Errorf("Expected %v, get: %v\n", cc.result, values)
		}
	}
}
