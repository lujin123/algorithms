package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPoisonedDuration(t *testing.T) {
	assert.EqualValues(t, 4, findPoisonedDuration([]int{1, 4}, 2))
	assert.EqualValues(t, 3, findPoisonedDuration([]int{1, 2}, 2))
	assert.EqualValues(t, 12, findPoisonedDuration([]int{1, 3, 5}, 8))
}
