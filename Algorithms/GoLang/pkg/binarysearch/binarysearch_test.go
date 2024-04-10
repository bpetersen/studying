package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func intComparer(a int, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

func stringComparer(a string, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

func TestBinarySearchInts(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	actual, found := BinarySearch[int](intComparer, vals, 3)

	assert.True(t, found)
	assert.Equal(t, 2, actual)
}

func TestBinarySearchStrings(t *testing.T) {
	vals := []string{"A", "B", "C", "D", "E"}
	actual, found := BinarySearch[string](stringComparer, vals, "B")

	assert.True(t, found)
	assert.Equal(t, 1, actual)
}

func TestBinarySearchFront(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	actual, found := BinarySearch[int](intComparer, vals, 1)

	assert.True(t, found)
	assert.Equal(t, 0, actual)
}

func TestBinarySearchEnd(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	actual, found := BinarySearch[int](intComparer, vals, 9)

	assert.True(t, found)
	assert.Equal(t, 8, actual)
}

func TestBinarySearchBeyondEnd(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	actual, found := BinarySearch[int](intComparer, vals, 10)

	assert.False(t, found)
	assert.Equal(t, 9, actual)
}

func TestBinarySearchBeyondBeginning(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	actual, found := BinarySearch[int](intComparer, vals, 0)

	assert.False(t, found)
	assert.Equal(t, 0, actual)
}
