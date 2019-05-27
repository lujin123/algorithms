package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	assert.Equal(t, "A", convert("A", 1))
	assert.Equal(t, "LCIRETOESIIGEDHN", convert("LEETCODEISHIRING", 3))
	assert.Equal(t, "LDREOEIIECIHNTSG", convert("LEETCODEISHIRING", 4))
}
