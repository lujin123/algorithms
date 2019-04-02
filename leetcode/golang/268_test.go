package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubmissions(t *testing.T) {
	assert.EqualValues(t, 2, missingNumber([]int{3, 0, 1}))
	assert.EqualValues(t, 0, missingNumber([]int{3, 2, 1}))
	assert.EqualValues(t, 8, missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

func TestSubmissions2(t *testing.T) {
	assert.EqualValues(t, 2, missingNumber2([]int{3, 0, 1}))
	assert.EqualValues(t, 0, missingNumber2([]int{3, 2, 1}))
	assert.EqualValues(t, 8, missingNumber2([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

func TestSubmissions3(t *testing.T) {
	assert.EqualValues(t, 2, missingNumber3([]int{3, 0, 1}))
	assert.EqualValues(t, 0, missingNumber3([]int{3, 2, 1}))
	assert.EqualValues(t, 8, missingNumber3([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
