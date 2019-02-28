package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	assert.Equal(t, "0", multiply("0", "0"))
	assert.Equal(t, "0", multiply("2", "0"))
	assert.Equal(t, "6", multiply("2", "3"))
	assert.Equal(t, "1234444433211", multiply("9999", "123456789"))
}
