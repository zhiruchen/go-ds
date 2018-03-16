package basic

import "errors"

const maxStackSize = 100000

type Stack struct {
	elements *LinkedList
	count int32
}

func NewStack() *Stack {
	return &Stack{elements:NewLinkedList(), count:0}
}

func (s *Stack) Push(data interface{}) {
	if s.count >= maxStackSize {
		panic("stack elements overflow!")
	}

	s.elements.InsertToFront(data)
	s.count++
}

func (s *Stack) Pop() (interface{}, error) {
	if s.count == 0 {
		return nil, errors.New("no element on stack")
	}

	head := s.elements.GetHead()
	if head == nil {
		return nil, errors.New("no element on stack")
	}

	s.elements.RemoveHead()
	s.count--
	return head.GetData()
}

func (s *Stack) Peek() (interface{}, error) {
	if s.count == 0 {
		return nil, errors.New("no element on stack")
	}

	head := s.elements.GetHead()
	if head == nil {
		return nil, errors.New("no element on stack")
	}

	return head.GetData()
}

func (s *Stack) IsEmpty() bool {
	return s.count == 0
}

func (s *Stack) String() string {
	return s.elements.String()
}