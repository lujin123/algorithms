package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
	assert.Equal(t, 3, majorityElement([]int{3, 2, 3}))
	assert.Equal(t, 2, majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
func TestMajorityElement2(t *testing.T) {
	assert.Equal(t, 3, majorityElement2([]int{3, 2, 3}))
	assert.Equal(t, 2, majorityElement2([]int{2, 2, 1, 1, 1, 2, 2}))
}
