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
	suppliedCapacity := 9
	list := NewArrayList[int](suppliedCapacity)
	assert.NotEqualValues(t, DEFAULT_INITIAL_CAPACITY, suppliedCapacity, "Default value is the same as supplied capacity.  Change the supplied value for accurate assertion in test.")
	if list.size != 0 {
		t.Errorf("expected size 0, got %d", list.size)
	}
	if len(list.data) != suppliedCapacity {
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

	actual, err := list.Set(0, "B")

	assert.NoError(t, err)
	assert.Equal(t, "A", actual)
	assert.Equal(t, []string{"", "", "B"}, list.data)
}

func TestGetReturnsIndexOutOfRangeErrors(t *testing.T) {
	list := NewArrayList[string]()
	_, underErr := list.Get(-1)

	assert.ErrorContains(t, underErr, "Index out of range")

	_, overErr := list.Get(1)

	assert.ErrorContains(t, overErr, "Index out of range")
}

func TestGetReturnsExistingItem(t *testing.T) {
	list := ArrayList[string]{
		data: []string{"", "", "A"},
		size: 1,
		head: 2,
	}

	actual, err := list.Get(0)

	assert.NoError(t, err)
	assert.Equal(t, "A", actual)
}

func TestAddReturnsIndexOutOfRangeErrors(t *testing.T) {
	list := NewArrayList[string]()
	underErr := list.Add(-1, "A")
	assert.ErrorContains(t, underErr, "Index out of range")

	overErr := list.Add(1, "A")
	assert.ErrorContains(t, overErr, "Index out of range")
}

func TestAddBasic(t *testing.T) {
	list := ArrayList[string]{
		data: []string{"", "", ""},
		size: 0,
		head: 0,
	}

	err := list.Add(0, "A")
	assert.NoError(t, err)
	assert.Equal(t, list.size, 1)
	assert.Equal(t, []string{"A", "", ""}, list.data)
}

func TestAddResizes(t *testing.T) {
	list := ArrayList[string]{
		data: []string{"A", "B", "C"},
		size: 3,
		head: 0,
	}

	err := list.Add(3, "D")
	assert.NoError(t, err)
	assert.Equal(t, []string{"A", "B", "C", "D", "", ""}, list.data)
}

func TestAddPushesLeft(t *testing.T) {
	list := ArrayList[string]{
		data: []string{"A", "C", "D", "E", "_"},
		size: 4,
		head: 0,
	}

	err := list.Add(1, "B")
	assert.NoError(t, err)
	assert.Equal(t, []string{"B", "C", "D", "E", "A"}, list.data)
}

func TestAddPushesRight(t *testing.T) {
	list := ArrayList[string]{
		data: []string{"_", "A", "B", "D", "E"},
		size: 4,
		head: 1,
	}

	err := list.Add(2, "C")
	assert.NoError(t, err)
	assert.Equal(t, []string{"E", "A", "B", "C", "D"}, list.data)
}

func TestRemoveReturnsIndexOutOfRangeErrors(t *testing.T) {
	list := NewArrayList[string]()
	_, underErr := list.Remove(-1)
	assert.ErrorContains(t, underErr, "Index out of range")

	_, overErr := list.Remove(1)
	assert.ErrorContains(t, overErr, "Index out of range")
}

func TestRemoveBasic(t *testing.T) {
	list := ArrayList[string]{
		data: []string{"A", "B", "C", "D", "E"},
		size: 5,
		head: 0,
	}

	actual, err := list.Remove(4)
	assert.NoError(t, err)
	assert.Equal(t, "E", actual)
	assert.Equal(t, 4, list.size)
	assert.Equal(t, []string{"A", "B", "C", "D", ""}, list.data)
}

func TestRemoveRightSideDoesntChangeHead(t *testing.T) {
	list := ArrayList[string]{
		data: []string{"C", "D", "E", "A", "B"},
		size: 5,
		head: 3,
	}

	actual, err := list.Remove(3)
	assert.NoError(t, err)
	assert.Equal(t, "D", actual)
	assert.Equal(t, 4, list.size)
	assert.Equal(t, 3, list.head)
	assert.Equal(t, []string{"C", "E", "", "A", "B"}, list.data)
}

func TestRemoveLeftSideIncrementsHead(t *testing.T) {
	list := ArrayList[string]{
		data: []string{"C", "D", "E", "A", "B"},
		size: 5,
		head: 3,
	}

	actual, err := list.Remove(1)
	assert.NoError(t, err)
	assert.Equal(t, "B", actual)
	assert.Equal(t, 4, list.size)
	assert.Equal(t, 4, list.head)
	assert.Equal(t, []string{"C", "D", "E", "", "A"}, list.data)
}

func TestRemoveResizes(t *testing.T) {
	list := ArrayList[string]{
		data: []string{"A", "B", "C", "", "", ""},
		size: 3,
		head: 0,
	}

	actual, err := list.Remove(2)
	assert.NoError(t, err)
	assert.Equal(t, "C", actual)
	assert.Equal(t, 2, list.size)
	assert.Equal(t, 0, list.head)
	assert.Equal(t, []string{"A", "B", "", ""}, list.data)
}

func TestRealign(t *testing.T) {
	list := ArrayList[string]{
		data: []string{"", "A", "B", "C", ""},
		size: 3,
		head: 1,
	}
	newData := make([]string, 7)
	list.realign(newData)
	assert.Equal(t, 0, list.head)
	assert.Equal(t, []string{"A", "B", "C", "", "", "", ""}, list.data)
}

func TestRemoveResizeRealign(t *testing.T) {
	list := ArrayList[string]{
		data: []string{"", "A", "B", "C", "", ""},
		size: 3,
		head: 1,
	}

	actual, err := list.Remove(0)
	assert.NoError(t, err)
	assert.Equal(t, "A", actual)
	assert.Equal(t, 2, list.size)
	assert.Equal(t, 0, list.head)
	assert.Equal(t, []string{"B", "C", "", ""}, list.data)
}
