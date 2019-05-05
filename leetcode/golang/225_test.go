package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyStack(t *testing.T) {
	obj := NewMyStack()
	obj.Push(1)
	obj.Push(2)
	assert.Equal(t, 2, obj.Pop())
	assert.Equal(t, 1, obj.Top())
	assert.Equal(t, false, obj.Empty())
	assert.Equal(t, 1, obj.Pop())
	assert.Equal(t, true, obj.Empty())
}
