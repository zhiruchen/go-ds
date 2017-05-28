package ds

type Node struct {
	Val   int32
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

// Search search a key in tree
func Search(root *Node, key int32) *Node {
	if root == nil || root.Val == key {
		return root
	}

	if key < root.Val {
		return Search(root.Left, key)
	}

	return Search(root.Right, key)
}

// NewNode create a new tree node
func NewNode(key int32) *Node {
	return &Node{
		Val:   key,
		Left:  nil,
		Right: nil,
	}
}

// Insert insert a key to bin search tree
func Insert(root *Node, key int32) *Node {
	if root == nil {
		return NewNode(key)
	}

	if key < root.Val {
		root.Left = Insert(root.Left, key)
	} else if key > root.Val {
		root.Right = Insert(root.Right, key)
	}
	return root
}

// DeleteNode delete a node from tree
func DeleteNode(root *Node, key int32) *Node {
	if root == nil {
		return root
	}

	if key < root.Left.Val {
		root.Left = DeleteNode(root.Left, key)
	} else if key > root.Right.Val {
		root.Right = DeleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		t := root
		root = min(root.Right)
		root.Right = deleteMinNode(t.Right)
		root.Left = t.Left
	}
	return root
}

//
func deleteMinNode(node *Node) *Node {
	if node.Left == nil {
		return node.Right
	}
	node.Left = deleteMinNode(node.Left)
	return node
}

// min find the min node of bin tree
func min(node *Node) *Node {
	if node.Left == nil {
		return node
	}
	return min(node.Left)
}

// GetKeys return a slice of keys
func GetKeys(root *Node) []int32 {
	nodes := []*Node{root}
	keys := []int32{}
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		keys = append(keys, node.Val)
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
	}
	return keys
}
