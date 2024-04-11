package sort

import (
	"algorithms/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortNoOp(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := []int{1, 2, 3, 4}

	actual := BubbleSort[int](utils.IntComparer, nums)
	assert.Equal(t, expected, actual)
}

func TestBubbleSortReverseOrder(t *testing.T) {
	nums := []int{4, 3, 2, 1}
	expected := []int{1, 2, 3, 4}

	actual := BubbleSort[int](utils.IntComparer, nums)
	assert.Equal(t, expected, actual)
}

func TestBubbleSort(t *testing.T) {
	nums := []int{4, 3, 2, 1, 7, 8, 9, 5, 6}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	actual := BubbleSort[int](utils.IntComparer, nums)
	assert.Equal(t, expected, actual)
}
