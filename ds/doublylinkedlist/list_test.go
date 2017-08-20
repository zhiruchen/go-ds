package doublylinkedlist

import "testing"
import "github.com/stretchr/testify/assert"

func TestInit(t *testing.T) {
	l := new(DoublyLinkedList).Init()
	assert.Equal(t, uint64(0), l.Len())
	assert.Equal(t, (*Node)(nil), l.Head())
	assert.Equal(t, (*Node)(nil), l.Tail())
}

func TestInsertBegin(t *testing.T) {
	l := new(DoublyLinkedList).Init()

	l.InsertBegin(&Node{Val: 1})
	assert.Equal(t, uint64(1), l.Len())
	assert.NotNil(t, l.Head())
	assert.Equal(t, 1, l.Head().Val.(int))

	l.InsertBegin(&Node{Val: 2})
	assert.Equal(t, uint64(2), l.Len())
	assert.Equal(t, 2, l.Head().Val.(int))
	assert.Equal(t, 1, l.Tail().Val.(int))
}

func TestIInsertEnd(t *testing.T) {
	l := new(DoublyLinkedList).Init()

	l.InsertEnd(&Node{Val: "hello"})
	assert.Equal(t, uint64(1), l.Len())
	assert.Equal(t, "hello", l.Head().Val.(string))
	assert.Equal(t, "hello", l.Tail().Val.(string))
	l.InsertEnd(&Node{Val: "world"})
	l.InsertEnd(&Node{Val: "a"})
	l.InsertEnd(&Node{Val: "programmer"})
	assert.Equal(t, uint64(4), l.Len())
	assert.Equal(t, "programmer", l.Tail().Val.(string))

}

func TestRemove(t *testing.T) {
	l := new(DoublyLinkedList).Init()

	l.InsertBegin(&Node{Val: "hello"})
	l.InsertBegin(&Node{Val: "a"})
	e3 := l.InsertBegin(&Node{Val: "world"})
	assert.Equal(t, uint64(3), l.Len())
	e3 = l.Remove(e3)
	assert.Equal(t, "world", e3.Val.(string))
	assert.Equal(t, uint64(2), l.Len())
	assert.Equal(t, "a", l.Head().Val.(string))

	l.InsertEnd(&Node{Val: "May"})
	e4 := l.InsertEnd(&Node{Val: "Be"})
	e4 = l.Remove(e4)
	assert.Equal(t, "Be", e4.Val.(string))
	assert.Equal(t, uint64(3), l.Len())

}
