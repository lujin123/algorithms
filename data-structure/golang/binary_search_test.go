package dsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	assert.Equal(t, true, binarySearch([]int{1}, 1))
	assert.Equal(t, true, binarySearch([]int{1, 2}, 1))
	assert.Equal(t, true, binarySearch([]int{1, 2, 3}, 2))
	assert.Equal(t, true, binarySearch([]int{1, 2, 3}, 3))
	assert.Equal(t, true, binarySearch([]int{1, 2, 3}, 1))
	assert.Equal(t, false, binarySearch([]int{1, 2, 3}, 0))

	assert.Equal(t, 0, linearSearch2([]int{1, 2}, 1))
	assert.Equal(t, -1, linearSearch2([]int{1, 2}, -1))
}
