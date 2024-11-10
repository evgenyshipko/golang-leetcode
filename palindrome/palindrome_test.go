package palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindrome(t *testing.T) {
	assert.Equal(t, true, IsPalindrome(121))
	assert.Equal(t, false, IsPalindrome(-121))
	assert.Equal(t, true, IsPalindrome(101101))
	assert.Equal(t, true, IsPalindrome(1))
	assert.Equal(t, false, IsPalindrome(167854378326))
}
