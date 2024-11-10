package remove_duplicates

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestRemoveDuplicatesCheckLength(t *testing.T) {
	testData := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	assert.Equal(t, 5, RemoveDuplicates(testData))
}

func TestRemoveDuplicatesCheckEntryArray(t *testing.T) {
	testData := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	length := RemoveDuplicates(testData)
	assert.Equal(t, true, reflect.DeepEqual(testData[:length], []int{0, 1, 2, 3, 4}))
}
