package basic

import (
	"fmt"
	"strings"
)

type newLinkListNode struct {
	data Comparator
	next *newLinkListNode
}

type LinkedListV2 struct {
	head  *newLinkListNode
	tail  *newLinkListNode
	count int32
}

func NewLinkedListV2() *LinkedListV2 {
	return &LinkedListV2{head: nil, tail: nil, count: 0}
}

func (ll *LinkedListV2) AppendToTail(data Comparator) {
	node := &newLinkListNode{data: data, next: nil}
	ll.count++

	if ll.head == nil {
		ll.head = node
		ll.tail = ll.head
	} else {
		ll.tail.next = node
		ll.tail = node
	}
}

func (ll *LinkedListV2) MergeSort() {
	ll.head = mergeSort(ll.head)
}

func (ll *LinkedListV2) String() string {
	var nodes []string
	node := ll.head
	for node != nil {
		data := fmt.Sprintf("%v", node.data)
		nodes = append(nodes, data)
		node = node.next
	}

	return strings.Join(nodes, "->")
}

func mergeSort(node *newLinkListNode) *newLinkListNode {
	if node == nil || node.next == nil {
		return node
	}

	middle := getMiddleOfLinkedList(node)
	nextOfMiddle := middle.next

	middle.next = nil
	left := mergeSort(node)
	right := mergeSort(nextOfMiddle)

	return merge(left, right)
}

func merge(x, y *newLinkListNode) *newLinkListNode {
	var res *newLinkListNode
	if x == nil {
		return y
	}

	if y == nil {
		return x
	}

	if x.data.Compare(y.data) <= 0 {
		res = x
		res.next = merge(x.next, y)
	} else {
		res = y
		res.next = merge(x, y.next)
	}

	return res
}

func getMiddleOfLinkedList(head *newLinkListNode) *newLinkListNode {
	if head == nil {
		return head
	}

	fast, slow := head.next, head
	for fast != nil {
		fast = fast.next
		if fast != nil {
			fast = fast.next
			slow = slow.next
		}
	}

	return slow
}

func MergeTwoSortedLinkedList(l1, l2 *LinkedListV2) *LinkedListV2 {
	l := NewLinkedListV2()
	var cur = &newLinkListNode{}
	l.head = cur

	head1, head2 := l1.head, l2.head
	for head1 != nil && head2 != nil {
		if head1.data.Compare(head2.data) <= 0 {
			cur.next = head1
			head1 = head1.next
		} else {
			cur.next = head2
			head2 = head2.next
		}

		cur = cur.next
	}

	for head1 != nil {
		cur.next = head1
		cur = cur.next
		head1 = head1.next
	}

	for head2 != nil {
		cur.next = head2
		cur = cur.next
		head2 = head2.next
	}

	l.head = l.head.next
	l.tail = cur
	return l
}
