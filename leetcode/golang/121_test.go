package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	assert.EqualValues(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
	assert.EqualValues(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))
}

func TestMaxProfit2(t *testing.T) {
	assert.EqualValues(t, 5, maxProfit2([]int{7, 1, 5, 3, 6, 4}))
	assert.EqualValues(t, 0, maxProfit2([]int{7, 6, 4, 3, 1}))
}
