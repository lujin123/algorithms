package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckRecordI(t *testing.T) {
	assert.Equal(t, true, checkRecord("PPALLP"))
	assert.Equal(t, false, checkRecord("PPALLL"))
	assert.Equal(t, false, checkRecord("LAPPA"))
}
