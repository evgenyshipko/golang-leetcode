package index_of_the_first_occurrence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndexOfFirstOccurence(t *testing.T) {
	assert.Equal(t, 0, IndexFirstOccur("sadbutsad", "sad"))
	assert.Equal(t, 2, IndexFirstOccur("bydedlol", "ded"))
	assert.Equal(t, -1, IndexFirstOccur("leetcode", "leeto"))
	assert.Equal(t, -1, IndexFirstOccur("", "leeto"))
	assert.Equal(t, 4, IndexFirstOccur("mississippi", "issip"))
	assert.Equal(t, 0, IndexFirstOccur("a", "a"))
}
