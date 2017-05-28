package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindFirstRepeatedChar(t *testing.T) {
	assert.Equal(t, 'e', FindFirstRepeatedChar("geekforgeeks"))
	assert.Equal(t, 'l', FindFirstRepeatedChar("hello"))
	assert.Equal(t, int32(-1), FindFirstRepeatedChar("abcd"))

}
