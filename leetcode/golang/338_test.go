package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBits(t *testing.T) {
	assert.Equal(t, []int{0, 1, 1}, countBits(2))
	assert.Equal(t, []int{0, 1, 1, 2, 1, 2}, countBits(5))
}

func TestCountBits2(t *testing.T) {
	assert.Equal(t, []int{0, 1, 1}, countBits2(2))
	assert.Equal(t, []int{0, 1, 1, 2, 1, 2}, countBits2(5))
}
