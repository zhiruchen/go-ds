package doublylinkedlist

// Node represent a list node in doublylinkedlist
type Node struct {
	Val  interface{}
	prev *Node
	next *Node
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) Next() *Node {
	return n.next
}

// DoublyLinkedList 双向链表
type DoublyLinkedList struct {
	head  *Node
	tail  *Node
	count uint64
}

// Init 初始化一个空链表
func (l *DoublyLinkedList) Init() *DoublyLinkedList {
	l.count = 0
	l.head = nil
	l.tail = nil
	return l
}

// InsertBefore 在e之前插入节点
func (l *DoublyLinkedList) InsertBefore(e, node *Node) *Node {
	node.next = e

	if e.prev == nil {
		node.prev = nil
		l.head = node
	} else {
		node.prev = e.prev
		node.prev.next = node
	}
	l.count++
	if e.next == nil {
		l.tail = e
	}
	return node
}

// InsertBegin 在链表开头插入节点
func (l *DoublyLinkedList) InsertBegin(node *Node) *Node {
	if l.head == nil {
		l.head = node
		l.tail = node

		node.prev, node.next = nil, nil
		l.count++
		return node
	}

	return l.InsertBefore(l.head, node)
}

// InsertAfter 在e节点之后插入node节点
func (l *DoublyLinkedList) InsertAfter(e, node *Node) *Node {
	node.prev = e
	if e.next == nil {
		node.next = nil
		l.tail = node
	} else {
		node.next = e.next
		e.next.prev = node
	}

	e.next = node
	l.count++
	return node
}

// InsertEnd 在链表末尾插入节点
func (l *DoublyLinkedList) InsertEnd(node *Node) *Node {
	if l.tail == nil {
		return l.InsertBegin(node)
	}
	return l.InsertAfter(l.tail, node)
}

// Remove 删除节点
func (l *DoublyLinkedList) Remove(node *Node) *Node {
	if node.prev == nil {
		l.head = node.next
		node.next.prev = nil
		l.count--
		return node
	}

	if node.next == nil {
		l.tail = node.prev
		node.prev.next = nil
		l.count--
		return node
	}

	node.prev.next = node.next
	node.next.prev = node.prev
	l.count--
	return node
}

func (l *DoublyLinkedList) MoveToFront(node *Node) {

	if l.count == 1 || l.count == 0 {
		return
	}

	if node != nil {
		prevNode := node.prev
		if prevNode == nil {
			return
		}
		nextNode := node.next

		if nextNode == nil {
			prevNode.next = nil
		} else {
			prevNode.next = nextNode
			nextNode.prev = prevNode
		}

		l.head.prev = node
		node.next = l.head
		node.prev = nil

		l.head = node
	}
}

// Len 返回链表的长度
func (l *DoublyLinkedList) Len() uint64 {
	return l.count
}

// Head 返回头部节点
func (l *DoublyLinkedList) Head() *Node {
	return l.head
}

// Tail 返回尾部节点
func (l *DoublyLinkedList) Tail() *Node {
	return l.tail
}
