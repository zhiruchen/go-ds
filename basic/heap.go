package basic

type BinaryHeap struct {
	elements []Comparer
	count    int
}

func NewBinaryHeap() *BinaryHeap {
	return &BinaryHeap{elements: []Comparer{nil}}
}

func (h *BinaryHeap) Insert(val Comparer) {
	h.elements = append(h.elements, val)
	h.HeapifyUp()
	h.count++
}

func (h *BinaryHeap) getParentIndex(childIndex int) int {
	return childIndex / 2
}

func (h *BinaryHeap) getChildIndex(parentIndex int) (left int, right int) {
	return parentIndex * 2, (parentIndex * 2) + 1
}

func (h *BinaryHeap) getChild(parentIndex int) (left Comparer, right Comparer) {
	l, r := h.getChildIndex(parentIndex)
	if l > len(h.elements)-1 {
		return nil, nil
	}

	if r > len(h.elements)-1 {
		return h.elements[l], nil
	}

	return h.elements[l], h.elements[r]
}

func (h *BinaryHeap) getMaxChild(parentIndex int) (Comparer, int) {
	left, right := h.getChild(parentIndex)

	if left == nil && right == nil {
		return nil, -1
	}

	if right == nil {
		return left, parentIndex * 2
	}

	if left.Less(right) {
		return right, parentIndex*2 + 1
	}

	return left, parentIndex * 2
}

func (h *BinaryHeap) HeapifyUp() {
	i := len(h.elements) - 1
	for i > 1 {
		parentIndex := h.getParentIndex(i)
		if h.elements[parentIndex].Less(h.elements[i]) {
			h.elements[i], h.elements[parentIndex] = h.elements[parentIndex], h.elements[i]
			i = parentIndex
		} else {
			break
		}
	}
}

func (h *BinaryHeap) HeapifyDown() {
	i := 1

	for i < len(h.elements) {
		child, childIndex := h.getMaxChild(i)
		if child == nil {
			break
		}

		if !h.elements[i].Less(child) {
			break
		}

		h.elements[i], h.elements[childIndex] = h.elements[childIndex], h.elements[i]
		i = childIndex
	}
}

func (h *BinaryHeap) DeleteRoot() Comparer {
	if h.count == 0 || len(h.elements) <= 1 {
		return nil
	}

	root := h.elements[1]
	length := len(h.elements)
	if length == 2 {
		h.elements = []Comparer{nil}
	} else {
		lastIndex := length - 1
		h.elements[1] = h.elements[lastIndex]
		h.elements = append(h.elements[:lastIndex], h.elements[lastIndex+1:]...)
		h.HeapifyDown()
	}

	h.count--
	return root
}

func (h *BinaryHeap) Max() Comparer {
	if h.count == 0 {
		return nil
	}

	return h.elements[1]
}

func (h *BinaryHeap) Sort() []Comparer {
	i := 1
	var vals []Comparer
	length := len(h.elements)
	for i < length {
		vals = append(vals, h.DeleteRoot())
		i++
	}

	return vals
}

func (h *BinaryHeap) LevelOrder() []Comparer {
	var vals []Comparer
	for _, e := range h.elements[1:] {
		vals = append(vals, e)
	}

	return vals
}
