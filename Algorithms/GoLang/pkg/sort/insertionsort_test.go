package sort

import (
	"algorithms/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSortNoOp(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := []int{1, 2, 3, 4}

	actual := InsertionSort[int](utils.IntComparer, nums)
	assert.Equal(t, expected, actual)
}

func TestInsertionSortReverseOrder(t *testing.T) {
	nums := []int{4, 3, 2, 1}
	expected := []int{1, 2, 3, 4}

	actual := InsertionSort[int](utils.IntComparer, nums)
	assert.Equal(t, expected, actual)
}

func TestInsertionSort(t *testing.T) {
	nums := []int{4, 3, 2, 1, 7, 8, 9, 5, 6}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	actual := InsertionSort[int](utils.IntComparer, nums)
	assert.Equal(t, expected, actual)
}
