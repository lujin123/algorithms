package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubstring(t *testing.T) {
	assert.Equal(t, 0, longestSubstring("ababacb", 3))
	assert.Equal(t, 3, longestSubstring("bbaaacbd", 3))
	assert.Equal(t, 0, longestSubstring("", 0))
	assert.Equal(t, 0, longestSubstring("", 3))
	assert.Equal(t, 3, longestSubstring("aaa", 3))
	assert.Equal(t, 3, longestSubstring("aaabb", 3))
	assert.Equal(t, 5, longestSubstring("ababbc", 2))
}
