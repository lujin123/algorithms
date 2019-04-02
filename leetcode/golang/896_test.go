package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMonotonic(t *testing.T) {
	assert.EqualValues(t, true, isMonotonic([]int{1, 2, 2, 3}))
	assert.EqualValues(t, true, isMonotonic([]int{6, 5, 4, 4}))
	assert.EqualValues(t, false, isMonotonic([]int{1, 3, 2}))
	assert.EqualValues(t, true, isMonotonic([]int{1, 2, 4, 5}))
	assert.EqualValues(t, true, isMonotonic([]int{1, 1, 1}))
}
