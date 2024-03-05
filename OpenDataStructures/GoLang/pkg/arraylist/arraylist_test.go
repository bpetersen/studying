package arraylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArrayListWithDefaultCapacity(t *testing.T) {
	list := NewArrayList[int]()
	if list.size != 0 {
		t.Errorf("expected size 0, got %d", list.size)
	}
	if len(list.data) != DEFAULT_INITIAL_CAPACITY {
		t.Errorf("expected data capacity %d, got %d", DEFAULT_INITIAL_CAPACITY, len(list.data))
	}
}

func TestNewArrayListWithSuppliedCapacity(t *testing.T) {
	suppliedCapacity := 4
	list := NewArrayList[int](suppliedCapacity)
	assert.NotEqualValues(t, DEFAULT_INITIAL_CAPACITY, suppliedCapacity, "Default value is the same as supplied capacity.  Change the supplied value for accurate assertion in test.")
	if list.size != 0 {
		t.Errorf("expected size 0, got %d", list.size)
	}
	if len(list.data) != 4 {
		t.Errorf("expected data capacity %d, got %d", DEFAULT_INITIAL_CAPACITY, len(list.data))
	}
}

func TestSize(t *testing.T) {
	list := ArrayList[int]{
		size: 7,
	}

	assert.Equal(t, 7, list.Size())
}

func TestSetReturnsIndexOutOfRangeErrors(t *testing.T) {
	list := NewArrayList[string]()
	_, underErr := list.Set(-1, "A")

	assert.ErrorContains(t, underErr, "Index out of range")

	_, overErr := list.Set(1, "A")

	assert.ErrorContains(t, overErr, "Index out of range")
}

func TestSetReturnsExistingItem(t *testing.T) {
	list := ArrayList[string]{
		data: []string{"", "", "A"},
		size: 1,
		head: 2,
	}

	result, err := list.Set(0, "B")

	assert.NoError(t, err)
	assert.Equal(t, result, "A")
	assert.Equal(t, []string{"", "", "B"}, list.data)
}
