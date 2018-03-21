package basic

type BSTNode struct {
	data  Comparer
	left  *BSTNode
	right *BSTNode
}

/*

  BinarySearchTree https://en.wikipedia.org/wiki/Binary_search_tree#Definition
*/
type BinarySearchTree struct {
	root *BSTNode
}

func (bst *BinarySearchTree) Search(key Comparer) *BSTNode {
	return bstSearch(bst.root, key)
}

func bstSearch(node *BSTNode, val Comparer) *BSTNode {
	if node == nil || node.data.Equal(val) {
		return node
	}

	data := node.data
	if val.Less(data) {
		return bstSearch(node.left, val)
	}

	return bstSearch(node.right, val)
}

func (bst *BinarySearchTree) Insert(key Comparer) {
	bst.root = bstInsert(bst.root, key)
}

func bstInsert(node *BSTNode, val Comparer) *BSTNode {
	if node == nil {
		node = &BSTNode{data: val}
		return node
	}

	if val.Equal(node.data) {
		node.data = val
	} else if val.Less(node.data) {
		node.left = bstInsert(node.left, val)
	} else {
		node.right = bstInsert(node.right, val)
	}

	return node
}

func (bst *BinarySearchTree) Delete(key *BSTNode) {
	if bst.root == nil {
		return
	}

	bst.root = deleteBSTNode(bst.root, key)
}

func deleteBSTNode(node *BSTNode, key *BSTNode) *BSTNode {
	if key == nil {
		return node
	}

	if key.data.Equal(node.data) {
		t := node
		if t.left == nil {
			return t.right
		}
		if t.right == nil {
			return t.left
		}

		node = findMin(t.right)
		node.right = deleteMin(t.right)
		node.left = t.left

	} else if key.data.Less(node.data) {
		node.left = deleteBSTNode(node.left, key)
	} else {
		node.right = deleteBSTNode(node.right, key)
	}

	return node
}

func (bst *BinarySearchTree) FindMin() *BSTNode {
	return findMin(bst.root)
}

func findMin(node *BSTNode) *BSTNode {
	if node == nil {
		return nil
	}

	for node.left != nil {
		node = node.left
	}

	return node
}

func (bst *BinarySearchTree) deleteMin() {
	if bst.root == nil {
		return
	}

	bst.root = deleteMin(bst.root)
}

func deleteMin(node *BSTNode) *BSTNode {
	if node.left == nil {
		return node.right
	}

	node.left = deleteMin(node.left)
	return node
}

func (bst *BinarySearchTree) LevelTraverse() ([]Comparer, error) {
	var sl []Comparer

	q := NewQueue()

	if bst.root != nil {
		q.Enqueue(bst.root)
	}

	for !q.IsEmpty() {
		node, err := q.Dequeue()
		if err != nil {
			return nil, err
		}

		nd := node.(*BSTNode)
		sl = append(sl, nd.data)

		if nd.left != nil {
			q.Enqueue(nd.left)
		}

		if nd.right != nil {
			q.Enqueue(nd.right)
		}
	}

	return sl, nil
}
