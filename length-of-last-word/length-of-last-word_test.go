package length_of_last_word

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {

	assert.Equal(t, 0, lengthOfLastWord(""))
	assert.Equal(t, 0, lengthOfLastWord(" "))
	assert.Equal(t, 0, lengthOfLastWord("    "))
	assert.Equal(t, 5, lengthOfLastWord("Hello World"))
	assert.Equal(t, 6, lengthOfLastWord(" рп1 Worldd   "))
	assert.Equal(t, 3, lengthOfLastWord(" 12 23 344 "))

	assert.Equal(t, 1, lengthOfLastWord(" a"))
	assert.Equal(t, 1, lengthOfLastWord("a"))
	assert.Equal(t, 1, lengthOfLastWord("a "))

	assert.Equal(t, 3, lengthOfLastWord("day"))

}
