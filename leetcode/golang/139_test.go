package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreak(t *testing.T) {
	assert.Equal(t, true, wordBreak("leetcode", []string{"leet", "code"}))
	assert.Equal(t, true, wordBreak("applepenapple", []string{"apple", "pen"}))
	assert.Equal(t, false, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

func TestWordBreak2(t *testing.T) {
	assert.Equal(t, true, wordBreak2("leetcode", []string{"leet", "code"}))
	assert.Equal(t, true, wordBreak2("applepenapple", []string{"apple", "pen"}))
	assert.Equal(t, false, wordBreak2("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
