package basic

import (
	"fmt"
	"strings"
	"errors"
)

type linkListNode struct {
	data interface{}
	next *linkListNode
}

func (n *linkListNode) GetData() (interface{}, error) {
	if n == nil {
		return nil, errors.New("node is nil")
	}

	return n.data, nil
}

type LinkedList struct {
	head *linkListNode
	tail *linkListNode
	count int32
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head:nil, tail: nil, count:0}
}

func (ll *LinkedList) InsertToFront(data interface{}) {
	node := &linkListNode{data:data, next:nil}
	ll.count++
	if ll.head == nil {
		ll.head = node
		ll.tail = ll.head
	} else {
		node.next = ll.head
		ll.head = node
	}
}

func (ll *LinkedList) AppendToTail(data interface{}) {
	node := &linkListNode{data:data, next:nil}
	ll.count++

	if ll.head == nil {
		ll.head = node
		ll.tail = ll.head
	} else {
		ll.tail.next = node
		ll.tail = node
	}
}

func (ll *LinkedList) DeleteNode(data interface{}) {
	headTemp := ll.head

	if headTemp == nil {
		return
	}

	if headTemp != nil && headTemp.data == data {
		ll.head = ll.head.next
		if ll.head == nil {
			ll.tail = nil
		}

		ll.count--
		return
	}

	for headTemp.next != nil {
		if headTemp.next.data == data {
			break
		}

		headTemp = headTemp.next
	}

	// delete headTemp.next
	if headTemp.next != nil {
		headTemp.next = headTemp.next.next

		if headTemp.next == nil {
			ll.tail = headTemp
		}
	}
}

func (ll *LinkedList) RemoveHead() {
	if ll.head == nil {
		return
	}

	ll.head = ll.head.next
	ll.count--
}

func (ll *LinkedList) GetHead() *linkListNode {
	return ll.head
}

func (ll *LinkedList) PrintList() {
	node := ll.head
	for node != nil {
		if node.next != nil {
			fmt.Printf("%v->", node.data)
		} else {
			fmt.Printf("%v\n", node.data)
		}

		node = node.next
	}
}

func (ll *LinkedList) String() string {
	var nodes []string
	node := ll.head
	for node != nil {
		data := fmt.Sprintf("%v", node.data)
		nodes = append(nodes, data)
		node = node.next
	}

	return strings.Join(nodes, "->")
}
