package dsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkiplist(t *testing.T) {
	sl := NewSkiplist()
	nums := []int{10, 12, 9, 7, 1, 19}
	for _, num := range nums {
		sl.Insert(num)
	}
	assert.Equal(t, true, sl.Find(1))
	assert.Equal(t, false, sl.Find(11))

	sl.Delete(10)
	sl.Delete(12)
	sl.Delete(7)
	sl.Delete(1)
	sl.Delete(9)
	sl.Delete(19)
	assert.Equal(t, 0, len(sl.levels))
}
