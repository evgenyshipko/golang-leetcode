package palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindrome(t *testing.T) {
	assert.Equal(t, IsPalindrome(121), true)
	assert.Equal(t, IsPalindrome(-121), false)
	assert.Equal(t, IsPalindrome(101101), true)
	assert.Equal(t, IsPalindrome(1), true)
	assert.Equal(t, IsPalindrome(167854378326), false)
}
