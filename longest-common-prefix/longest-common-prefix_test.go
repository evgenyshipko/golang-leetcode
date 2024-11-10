package longest_common_prefix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	assert.Equal(t, "flow", LongestCommonPrefix([]string{"flower", "flow", "flowht"}))
	assert.Equal(t, "", LongestCommonPrefix([]string{"dog", "racecar", "car"}))
	assert.Equal(t, "", LongestCommonPrefix([]string{"", "racecar", "car"}))
}
