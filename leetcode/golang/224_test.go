package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	assert.Equal(t, 2, calculate("1 + 1"))
	assert.Equal(t, 3, calculate(" 2-1 + 2 "))
	assert.Equal(t, 23, calculate("(1+(4+5+2)-3)+(6+8)"))
}
