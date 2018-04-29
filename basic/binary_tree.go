package basic

type BinTreeNode struct {
	data  interface{}
	left  *BinTreeNode
	right *BinTreeNode
}

type BinaryTree struct {
	root *BinTreeNode
}

func NewBinaryTree(root *BinTreeNode) *BinaryTree {
	return &BinaryTree{root: root}
}

// 前序遍历  root -> left -> right
func (b *BinaryTree) PreOrder() []interface{} {
	var sl []interface{}
	if b.root != nil {
		preOrder(b.root, &sl)
	}

	return sl
}

func preOrder(node *BinTreeNode, sl *[]interface{}) {
	if node == nil {
		return
	}

	*sl = append(*sl, node.data)

	if node.left != nil {
		preOrder(node.left, sl)
	}

	if node.right != nil {
		preOrder(node.right, sl)
	}
}

// 中序遍历 left -> root -> right
func (b *BinaryTree) InOrder() []interface{} {
	var sl []interface{}

	if b.root != nil {
		inOrder(b.root, &sl)
	}

	return sl
}

func inOrder(node *BinTreeNode, sl *[]interface{}) {
	if node == nil {
		return
	}

	if node.left != nil {
		inOrder(node.left, sl)
	}

	*sl = append(*sl, node.data)

	if node.right != nil {
		inOrder(node.right, sl)
	}
}

// 后续遍历 left -> right -> root
func (b *BinaryTree) PostOrder() []interface{} {
	var sl []interface{}

	if b.root != nil {
		postOrder(b.root, &sl)
	}

	return sl
}

func postOrder(node *BinTreeNode, sl *[]interface{}) {
	if node == nil {
		return
	}

	if node.left != nil {
		postOrder(node.left, sl)
	}

	if node.right != nil {
		postOrder(node.right, sl)
	}

	*sl = append(*sl, node.data)
}

func (b *BinaryTree) LevelOrder() ([]interface{}, error) {
	var sl []interface{}

	q := NewQueue()

	if b.root != nil {
		q.Enqueue(b.root)
	}

	for !q.IsEmpty() {
		node, err := q.Dequeue()
		if err != nil {
			return nil, err
		}

		nd := node.(*BinTreeNode)
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

func (b *BinaryTree) LowestCommonAncestor(n1, n2 interface{}) *BinTreeNode {
	return lowestCommonAncestor(b.root, n1, n2)
}

func lowestCommonAncestor(node *BinTreeNode, n1, n2 interface{}) *BinTreeNode {
	if node == nil {
		return nil
	}

	if node.data == n1 || node.data == n2 {
		return node
	}

	leftLCA := lowestCommonAncestor(node.left, n1, n2)
	rightLCA := lowestCommonAncestor(node.right, n1, n2)

	if leftLCA != nil && rightLCA != nil {
		return node
	}

	if leftLCA != nil {
		return leftLCA
	}

	return rightLCA
}

// IsSubTree 检查t2是否为当前二叉树的子树
func (b *BinaryTree) IsSubTree(t2 *BinaryTree) bool {
	return isSubTree(b.root, t2.root)
}

// 检查t2是否是t1的子树
func isSubTree(t1, t2 *BinTreeNode) bool {
	if t2 == nil {
		return true
	}

	if t1 == nil {
		return false
	}

	if isIdentical(t1, t2) {
		return true
	}

	return isSubTree(t1.left, t2) || isSubTree(t1.right, t2)
}

func isIdentical(r1, r2 *BinTreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}

	if r1 == nil || r2 == nil {
		return false
	}

	return r1.data == r2.data && isIdentical(r1.left, r2.left) && isIdentical(r1.right, r2.right)
}
