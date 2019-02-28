package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInclusion(t *testing.T) {
	assert.Equal(t, true, checkInclusion("ab", "eidbaooo"))
	assert.Equal(t, false, checkInclusion("ab", "eidboaoo"))
	assert.Equal(t, true, checkInclusion("", ""))
	assert.Equal(t, false, checkInclusion("abc", "ab"))
}

func TestCheckInclusion2(t *testing.T) {
	assert.Equal(t, true, checkInclusion2("ab", "eidbaooo"))
	assert.Equal(t, false, checkInclusion2("ab", "eidboaoo"))
	assert.Equal(t, true, checkInclusion2("", ""))
	assert.Equal(t, false, checkInclusion2("abc", "ab"))
}
