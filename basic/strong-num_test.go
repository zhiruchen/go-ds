package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsStrongNumber(t *testing.T) {
	assert.Equal(t, false, IsStrongNumber(123))
	assert.Equal(t, true, IsStrongNumber(145))
	assert.Equal(t, false, IsStrongNumber(123))
	assert.Equal(t, false, IsStrongNumber(123))
	assert.Equal(t, false, IsStrongNumber(123))
	assert.Equal(t, false, IsStrongNumber(123))
}
