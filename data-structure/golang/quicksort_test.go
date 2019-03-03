package dsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuicksort(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, quicksort([]int{3, 2, 1}))
	assert.Equal(t, []int{-1, -1, 0, 1, 2, 3, 3}, quicksort([]int{3, 2, 1, -1, -1, 0, 3}))
}
