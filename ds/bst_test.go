package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertNode(t *testing.T) {
	root := Insert(nil, 3)
	root = Insert(root, 2)
	root = Insert(root, 1)
	root = Insert(root, 4)
	root = Insert(root, 5)
	root = Insert(root, 6)

	keys := GetKeys(root)
	assert.Equal(t, []int32{3, 2, 4, 1, 5, 6}, keys)
}

func TestDeleteNode(t *testing.T) {
	root := Insert(nil, 3)
	root = Insert(root, 2)
	root = Insert(root, 1)
	root = Insert(root, 4)
	root = Insert(root, 5)
	root = Insert(root, 6)

	root = DeleteNode(root, 3)

	keys := GetKeys(root)
	assert.Equal(t, []int32{4, 2, 5, 1, 6}, keys)
}
