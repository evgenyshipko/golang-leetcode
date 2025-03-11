package __no_leet_code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeArrays(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, mergeSlices([]int{1, 2}, []int{3, 4, 5}, []int{6}))
}
