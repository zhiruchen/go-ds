package advanced

type TSTNode struct {
	val              rune
	endOfStr         bool
	left, mid, right *TSTNode
}

// TST Ternary Search Tree
type TST struct {
	root *TSTNode
}

func NewTST() *TST {
	return &TST{
		root: nil,
	}
}

func (t *TST) Put(key string) {
	t.root = put(t.root, key, 0)
}

func put(node *TSTNode, key string, d int) *TSTNode {
	val := rune(key[d])

	if node == nil {
		node = &TSTNode{
			val: val,
		}
	}

	if val < node.val {
		node.left = put(node.left, key, d)
	} else if val > node.val {
		node.right = put(node.right, key, d)
	} else {
		if d < len(key)-1 {
			node.mid = put(node.mid, key, d+1)
		} else {
			node.val = val
			node.endOfStr = true
		}
	}

	return node
}

func (t *TST) Contains(key string) bool {
	node := t.get(t.root, key, 0)
	if node == nil {
		return false
	}

	return node.endOfStr
}

func (t *TST) get(node *TSTNode, key string, d int) *TSTNode {
	if node == nil {
		return nil
	}

	v := rune(key[d])

	if v < node.val {
		return t.get(node.left, key, d)
	}

	if v > node.val {
		return t.get(node.right, key, d)
	}

	if d < len(key)-1 {
		return t.get(node.mid, key, d+1)
	}

	return node
}
