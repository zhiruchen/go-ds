package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPut(t *testing.T) {
	h := NewMap()
	assert.Equal(t, uint64(0), h.Count)
	assert.Equal(t, uint64(16), h.Size)
	h.Put("a", 1)
	h.Put("b", 2)
	h.Put("c", 3)

	assert.Equal(t, uint64(3), h.Count)
	assert.Equal(t, uint64(16), h.Size)
	assert.Equal(t, .1875, h.LoadFactor)

	assert.Equal(t, 1, h.Get("a").(int))
	assert.Equal(t, 2, h.Get("b").(int))
	assert.Equal(t, 3, h.Get("c").(int))
	assert.Equal(t, nil, h.Get("x"))
}

func TestDelete(t *testing.T) {
	m := NewMap()
	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)
	m.Put("Hello World", 42)
	m.Put("Good", "Bad")

	v := m.Delete("A")
	assert.Equal(t, nil, v)
	assert.Equal(t, uint64(5), m.Count)

	v = m.Delete("b")
	assert.Equal(t, 2, v.(int))
	assert.Equal(t, uint64(4), m.Count)
}
