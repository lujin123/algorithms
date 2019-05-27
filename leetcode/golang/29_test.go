package leetcode

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	assert.Equal(t, 3, divide(10, 3))
	assert.Equal(t, -2, divide(7, -3))
	assert.Equal(t, math.MaxInt32, divide(-2147483648, -1))
}
