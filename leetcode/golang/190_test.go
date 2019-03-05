package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBits(t *testing.T) {
	assert.Equal(t, uint32(964176192), reverseBits(43261596))
	assert.Equal(t, uint32(3221225471), reverseBits(4294967293))
}
