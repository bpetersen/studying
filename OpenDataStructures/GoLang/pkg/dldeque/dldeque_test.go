package dldeque

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	deque := NewDLDeque[string]()

	deque.AddFirst("C")
	deque.AddLast("D")
	deque.AddFirst("B")
	deque.AddLast("E")
	deque.AddFirst("A")
	deque.AddLast("F")

	var actual string

	assert.Equal(t, 6, deque.Size())

	actual, _ = deque.RemoveFirst()
	assert.Equal(t, "A", actual)
	assert.Equal(t, 5, deque.Size())

	actual, _ = deque.RemoveFirst()
	assert.Equal(t, "B", actual)
	assert.Equal(t, 4, deque.Size())

	actual, _ = deque.RemoveFirst()
	assert.Equal(t, "C", actual)
	assert.Equal(t, 3, deque.Size())

	actual, _ = deque.RemoveLast()
	assert.Equal(t, "F", actual)
	assert.Equal(t, 2, deque.Size())

	actual, _ = deque.RemoveLast()
	assert.Equal(t, "E", actual)
	assert.Equal(t, 1, deque.Size())

	actual, _ = deque.RemoveLast()
	assert.Equal(t, "D", actual)
	assert.Equal(t, 0, deque.Size())

	var err error
	_, err = deque.RemoveLast()
	assert.ErrorContains(t, err, "DLDeque is empty")

	_, err = deque.RemoveFirst()
	assert.ErrorContains(t, err, "DLDeque is empty")
}
